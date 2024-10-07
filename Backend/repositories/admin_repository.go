package repositories

import (
	"GOCOM/models"
	"GOCOM/utils/errors"

	"gorm.io/gorm"
)

// AdminUserRepositoryer 定义了管理员用户仓库的接口
type AdminUserRepositoryer interface {
	GetAllUsersWithRoles(page, pageSize int, departmentID *uint, email, name string, roleIDs []uint) ([]models.AdminUser, int64, error)
	CreateAdminUser(user *models.AdminUser) error
	UpdateAdminUser(user *models.AdminUser) error
	DeleteAdminUser(userID uint) error
	FindUserByID(userID uint) (*models.AdminUser, error)
	FindUserByEmail(email string) (*models.AdminUser, error)
	FindUserByPhone(phone string) (*models.AdminUser, error)
	getAllDepartmentIDs(departmentID uint) ([]uint, error)
	GetUserByEmail(email string) (*models.AdminUser, error)
	GetUserByID(id uint) (*models.AdminUser, error)
}

// AdminUserRepository 实现了 AdminUserRepositoryer 接口
type AdminUserRepository struct {
	DB *gorm.DB
}

// NewAdminUserRepository 创建一个新的 AdminUserRepository 实例
func NewAdminUserRepository(db *gorm.DB) AdminUserRepositoryer {
	return &AdminUserRepository{DB: db}
}

// GetAllUsersWithRoles 获取所有用户及其角色
func (repo *AdminUserRepository) GetAllUsersWithRoles(page, pageSize int, departmentID *uint, email, name string, roleIDs []uint) ([]models.AdminUser, int64, error) {
	var users []models.AdminUser
	var total int64

	query := repo.DB.Model(&models.AdminUser{}).
		Preload("Roles").
		Preload("Department")

	if departmentID != nil {
		departmentIDs, err := repo.getAllDepartmentIDs(*departmentID)
		if err != nil {
			return nil, 0, err
		}
		query = query.Where("department_id IN ?", departmentIDs)
	}

	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if len(roleIDs) > 0 {
		query = query.Distinct().
			Joins("JOIN user_roles ON user_roles.user_id = admin_users.id").
			Where("user_roles.role_id IN ?", roleIDs)
	}

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err = query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CreateAdminUser 创建admin用户
func (repo *AdminUserRepository) CreateAdminUser(user *models.AdminUser) error {
	return repo.DB.Create(user).Error
}

// UpdateAdminUser 更新admin用户
func (repo *AdminUserRepository) UpdateAdminUser(user *models.AdminUser) error {
	return repo.DB.Model(user).Omit("created_at").Updates(user).Error
}

// DeleteAdminUser 删除admin用户
func (repo *AdminUserRepository) DeleteAdminUser(userID uint) error {
	return repo.DB.Delete(&models.AdminUser{}, userID).Error
}

// FindUserByID 根据ID查找用户
func (repo *AdminUserRepository) FindUserByID(userID uint) (*models.AdminUser, error) {
	var user models.AdminUser
	err := repo.DB.Preload("Roles").First(&user, userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("用户不存在", err)
		}
		return nil, errors.NewInternalServerError("查找用户失败", err)
	}
	return &user, nil
}

// FindUserByEmail 查找具有指定邮箱的用户
func (repo *AdminUserRepository) FindUserByEmail(email string) (*models.AdminUser, error) {
	var user models.AdminUser
	err := repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.NewInternalServerError("查找用户失败", err)
	}
	return &user, nil
}

// FindUserByPhone 查找具有指定手机号的用户
func (repo *AdminUserRepository) FindUserByPhone(phone string) (*models.AdminUser, error) {
	var user models.AdminUser
	err := repo.DB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.NewInternalServerError("查找用户失败", err)
	}
	return &user, nil
}

// getAllDepartmentIDs 获取给定部门ID及其所有子部门的ID
func (repo *AdminUserRepository) getAllDepartmentIDs(departmentID uint) ([]uint, error) {
	var departments []models.Department
	var departmentIDs []uint

	// 获取所有部门
	if err := repo.DB.Find(&departments).Error; err != nil {
		return nil, err
	}

	// 构建部门树
	departmentMap := make(map[uint][]models.Department)
	for _, dept := range departments {
		parentID := uint(0)
		if dept.ParentID != nil {
			parentID = *dept.ParentID
		}
		departmentMap[parentID] = append(departmentMap[parentID], dept)
	}

	// 递归获取所有子部门ID
	var getDepartmentIDs func(id uint)
	getDepartmentIDs = func(id uint) {
		departmentIDs = append(departmentIDs, id)
		for _, child := range departmentMap[id] {
			getDepartmentIDs(child.ID)
		}
	}

	getDepartmentIDs(departmentID)

	return departmentIDs, nil
}

func (repo *AdminUserRepository) GetUserByEmail(email string) (*models.AdminUser, error) {
	var user models.AdminUser
	err := repo.DB.Preload("Roles").Preload("Department").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo *AdminUserRepository) GetUserByID(id uint) (*models.AdminUser, error) {
	var user models.AdminUser
	err := repo.DB.Preload("Role").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
