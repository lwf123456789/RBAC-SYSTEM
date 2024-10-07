package repositories

import (
	"GOCOM/models"
	"GOCOM/utils/errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// MenuRepository 接口定义
type IMenuRepository interface {
	Create(menu *models.Menu) error
	GetAll(page, pageSize int) ([]models.Menu, int64, error)
	GetAllWithRoles(page, pageSize int) ([]models.Menu, int64, error)
	GetByID(id uint) (*models.Menu, error)
	Update(menu *models.Menu) error
	Delete(id uint) error
	GetUserMenus(userID uint) ([]models.Menu, error)
	AddMenuPermissions(menuID uint, permissions []models.Permission) error
	GetMenuPermissions(menuID uint) ([]models.Permission, error)
}

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) IMenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) Create(menu *models.Menu) error {
	return r.db.Create(menu).Error
}

func (r *MenuRepository) GetAll(page, pageSize int) ([]models.Menu, int64, error) {
	var menus []models.Menu
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&models.Menu{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(pageSize).Find(&menus).Error
	if err != nil {
		return nil, 0, err
	}

	return menus, total, nil
}

func (r *MenuRepository) GetAllWithRoles(page, pageSize int) ([]models.Menu, int64, error) {
	var menus []models.Menu
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&models.Menu{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Roles").Offset(offset).Limit(pageSize).Find(&menus).Error
	if err != nil {
		return nil, 0, err
	}

	return menus, total, nil
}

func (r *MenuRepository) GetByID(id uint) (*models.Menu, error) {
	var menu models.Menu
	err := r.db.First(&menu, id).Error
	return &menu, err
}

func (r *MenuRepository) Update(menu *models.Menu) error {
	return r.db.Omit("created_at").Save(menu).Error
}

func (r *MenuRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 首先删除 role_menus 表中的相关记录
		if err := tx.Where("menu_id = ?", id).Delete(&models.RoleMenu{}).Error; err != nil {
			fmt.Printf("删除role_menus记录时出错: %v\n", err)
			return err
		}

		// 然后删除菜单
		if err := tx.Delete(&models.Menu{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *MenuRepository) GetUserMenus(userID uint) ([]models.Menu, error) {
	var user models.AdminUser
	if err := r.db.Preload("Roles").First(&user, userID).Error; err != nil {
		return nil, err
	}

	var menus []models.Menu
	var query *gorm.DB

	query = r.db.Preload("PermissionItems")

	if !user.IsSuperAdmin() {
		query = query.
			Joins("JOIN role_menus ON menus.id = role_menus.menu_id").
			Joins("JOIN user_roles ON role_menus.role_id = user_roles.role_id").
			Where("user_roles.user_id = ?", userID)
	}

	err := query.Find(&menus).Error
	if err != nil {
		return nil, err
	}

	for i := range menus {
		var permissions []string
		if user.IsSuperAdmin() {
			for _, perm := range menus[i].PermissionItems {
				permissions = append(permissions, perm.Code)
			}
		} else {
			err := r.db.Model(&models.Permission{}).
				Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
				Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
				Where("user_roles.user_id = ? AND permissions.menu_id = ?", userID, menus[i].ID).
				Pluck("permissions.code", &permissions).Error
			if err != nil {
				return nil, err
			}
		}
		menus[i].SetPermissions(permissions)
	}

	return menus, nil
}

func (r *MenuRepository) AddMenuPermissions(menuID uint, newPermissions []models.Permission) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var existingPermissions []models.Permission
		if err := tx.Preload("Roles").Where("menu_id = ?", menuID).Find(&existingPermissions).Error; err != nil {
			return err
		}

		existingPermMap := make(map[string]models.Permission)
		for _, p := range existingPermissions {
			existingPermMap[p.Code] = p
		}

		var permissionsToCreate, permissionsToUpdate []models.Permission
		var updatedPermissionCodes []string
		rolePermissionsToAddMap := make(map[string]struct{})
		rolePermissionsToDeleteMap := make(map[string]struct{})

		for _, newPerm := range newPermissions {
			if existingPerm, exists := existingPermMap[newPerm.Code]; exists {
				existingPerm.Description = newPerm.Description
				permissionsToUpdate = append(permissionsToUpdate, existingPerm)
				delete(existingPermMap, newPerm.Code)

				existingRoleIDs := make(map[uint]bool)
				for _, role := range existingPerm.Roles {
					existingRoleIDs[role.ID] = true
				}

				for _, newRoleID := range newPerm.RoleIDs {
					if !existingRoleIDs[newRoleID] {
						key := fmt.Sprintf("%d_%d", newRoleID, existingPerm.ID)
						rolePermissionsToAddMap[key] = struct{}{}
					}
					delete(existingRoleIDs, newRoleID)
				}

				for roleID := range existingRoleIDs {
					key := fmt.Sprintf("%d_%d", roleID, existingPerm.ID)
					rolePermissionsToDeleteMap[key] = struct{}{}
				}
			} else {
				newPerm.MenuID = &menuID
				permissionsToCreate = append(permissionsToCreate, newPerm)
			}
			updatedPermissionCodes = append(updatedPermissionCodes, newPerm.Code)
		}

		if len(permissionsToCreate) > 0 {
			if err := tx.Create(&permissionsToCreate).Error; err != nil {
				return err
			}
			for _, newPerm := range permissionsToCreate {
				for _, roleID := range newPerm.RoleIDs {
					key := fmt.Sprintf("%d_%d", roleID, newPerm.ID)
					rolePermissionsToAddMap[key] = struct{}{}
				}
			}
		}

		if len(permissionsToUpdate) > 0 {
			if err := tx.Save(&permissionsToUpdate).Error; err != nil {
				return err
			}
		}

		permissionsToDelete := make([]uint, 0, len(existingPermMap))
		for _, p := range existingPermMap {
			permissionsToDelete = append(permissionsToDelete, p.ID)
		}
		if len(permissionsToDelete) > 0 {
			if err := tx.Delete(&models.Permission{}, permissionsToDelete).Error; err != nil {
				return err
			}
		}

		if len(rolePermissionsToAddMap) > 0 {
			values := make([]string, 0, len(rolePermissionsToAddMap))
			args := make([]interface{}, 0, len(rolePermissionsToAddMap)*2)
			for key := range rolePermissionsToAddMap {
				var roleID, permID uint
				fmt.Sscanf(key, "%d_%d", &roleID, &permID)
				values = append(values, "(?, ?)")
				args = append(args, roleID, permID)
			}
			query := fmt.Sprintf("INSERT INTO role_permissions (role_id, permission_id) VALUES %s ON DUPLICATE KEY UPDATE role_id=VALUES(role_id)", strings.Join(values, ","))
			if err := tx.Exec(query, args...).Error; err != nil {
				return err
			}
		}

		if len(rolePermissionsToDeleteMap) > 0 {
			deleteQuery := "DELETE FROM role_permissions WHERE (role_id, permission_id) IN (?)"
			var deleteArgs [][]interface{}
			for key := range rolePermissionsToDeleteMap {
				var roleID, permID uint
				fmt.Sscanf(key, "%d_%d", &roleID, &permID)
				deleteArgs = append(deleteArgs, []interface{}{roleID, permID})
			}
			if err := tx.Exec(deleteQuery, deleteArgs).Error; err != nil {
				return err
			}
		}

		var menu models.Menu
		if err := tx.First(&menu, menuID).Error; err != nil {
			return err
		}
		if err := menu.SetPermissions(updatedPermissionCodes); err != nil {
			return err
		}
		return tx.Save(&menu).Error
	})
}

func (r *MenuRepository) GetMenuPermissions(menuID uint) ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.Where("menu_id = ?", menuID).Find(&permissions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("菜单未找到", err)
		}
		return nil, err
	}

	for i := range permissions {
		var roleIDs []uint
		err := r.db.Table("role_permissions").
			Where("permission_id = ?", permissions[i].ID).
			Pluck("role_id", &roleIDs).Error
		if err != nil {
			return nil, err
		}
		permissions[i].RoleIDs = roleIDs
	}

	return permissions, nil
}
