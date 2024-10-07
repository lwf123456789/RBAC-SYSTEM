// repositories/role_repository.go

package repositories

import (
	"GOCOM/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// IRoleRepository 接口定义
type IRoleRepository interface {
	CreateRole(role *models.Role) error
	GetAllRoles(page, pageSize int) ([]models.Role, int64, error)
	GetRoleByID(id uint) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
	AssignMenusToRole(roleID uint, menuIDs []uint) error
	GetAssignedMenuIDs(roleID uint) ([]uint, error)
}

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &RoleRepository{db}
}

func (r *RoleRepository) CreateRole(role *models.Role) error {
	// 开始事务
	tx := r.db.Begin()

	// 创建角色
	if err := tx.Create(role).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 在字典表中创建对应的记录
	dictionary := models.Dictionary{
		Type:        "ROLE_TYPE", // 类型设置为 "role"
		Label:       role.Name,
		Value:       fmt.Sprintf("%d", role.ID), // 将角色的 ID 转换为字符串作为值
		Description: role.Description,
	}

	if err := tx.Create(&dictionary).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

func (r *RoleRepository) GetAllRoles(page int, pageSize int) ([]models.Role, int64, error) {
	var roles []models.Role
	var total int64

	// 确保 page 和 pageSize 有效
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 查询角色总数
	err := r.db.Model(&models.Role{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询角色
	err = r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&roles).Error
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

func (r *RoleRepository) GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) UpdateRole(role *models.Role) error {
	if role.ID == 0 {
		return errors.New("缺少角色ID")
	}
	return r.db.Model(&role).Omit("created_at").Updates(role).Error
}

func (r *RoleRepository) DeleteRole(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除 user_roles 中间表的关联数据
		if err := tx.Where("role_id = ?", id).Delete(&models.UserRole{}).Error; err != nil {
			return err
		}

		// 删除角色
		if err := tx.Delete(&models.Role{}, id).Error; err != nil {
			return err
		}

		// 删除字典表中对应的记录
		if err := tx.Where("type = ? AND value = ?", "ROLE_TYPE", fmt.Sprintf("%d", id)).Delete(&models.Dictionary{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *RoleRepository) AssignMenusToRole(roleID uint, menuIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 获取所有菜单
		var allMenus []models.Menu
		if err := tx.Find(&allMenus).Error; err != nil {
			return err
		}

		// 创建菜单ID到菜单的映射
		menuMap := make(map[uint]*models.Menu)
		for i := range allMenus {
			menuMap[allMenus[i].ID] = &allMenus[i]
		}

		// 创建需要分配的菜单集合，包括父菜单
		menuSet := make(map[uint]struct{})
		for _, id := range menuIDs {
			menuSet[id] = struct{}{}
			// 添加所有父菜单
			for menu := menuMap[id]; menu != nil && menu.ParentID != nil; menu = menuMap[*menu.ParentID] {
				menuSet[*menu.ParentID] = struct{}{}
			}
		}

		// 将集合转换回切片
		completeMenuIDs := make([]uint, 0, len(menuSet))
		for id := range menuSet {
			completeMenuIDs = append(completeMenuIDs, id)
		}

		// 删除所有现有的角色-菜单关联
		if err := tx.Where("role_id = ?", roleID).Delete(&models.RoleMenu{}).Error; err != nil {
			return err
		}

		// 创建新的角色-菜单关联
		for _, menuID := range completeMenuIDs {
			if err := tx.Create(&models.RoleMenu{RoleID: roleID, MenuID: menuID}).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *RoleRepository) GetAssignedMenuIDs(roleID uint) ([]uint, error) {
	var menuIDs []uint
	err := r.db.Model(&models.RoleMenu{}).Where("role_id = ?", roleID).Pluck("menu_id", &menuIDs).Error
	return menuIDs, err
}

// 辅助函数：找出在 a 中但不在 b 中的元素
// func difference(a, b []uint) []uint {
// 	mb := make(map[uint]struct{}, len(b))
// 	for _, x := range b {
// 		mb[x] = struct{}{}
// 	}
// 	var diff []uint
// 	for _, x := range a {
// 		if _, found := mb[x]; !found {
// 			diff = append(diff, x)
// 		}
// 	}
// 	return diff
// }
