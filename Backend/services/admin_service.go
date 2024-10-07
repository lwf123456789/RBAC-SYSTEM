package services

import (
	"GOCOM/models"
	"GOCOM/repositories"
	"GOCOM/utils/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AdminUserServicer 定义了管理员用户服务的接口
type AdminUserServicer interface {
	GetAllUsersWithRoles(page, pageSize int, departmentID *uint, email, name string, roleIDs []uint) ([]models.AdminUser, int64, error)
	CreateAdminUser(user *models.AdminUser) error
	UpdateAdminUser(user *models.AdminUser) error
	DeleteAdminUser(userID uint) error
	GetAdminUserByID(userID uint) (*models.AdminUser, error)
	ChangePassword(userID uint, oldPassword, newPassword string) error
	GetUserInfo(userID uint) (*models.AdminUser, error)
	Login(email, password string) (string, *models.AdminUser, error)
}

// AdminUserService 实现了 AdminUserServicer 接口
type AdminUserService struct {
	repo repositories.AdminUserRepositoryer
}

// NewAdminUserService 创建一个新的 AdminUserService 实例
func NewAdminUserService(repo repositories.AdminUserRepositoryer) AdminUserServicer {
	return &AdminUserService{repo: repo}
}

// GetAllUsersWithRoles 获取所有用户及其角色，支持分页和部门过滤
func (s *AdminUserService) GetAllUsersWithRoles(page, pageSize int, departmentID *uint, email, name string, roleIDs []uint) ([]models.AdminUser, int64, error) {
	return s.repo.GetAllUsersWithRoles(page, pageSize, departmentID, email, name, roleIDs)
}

// CreateAdminUser 创建新的管理员用户
func (s *AdminUserService) CreateAdminUser(user *models.AdminUser) error {
	if err := s.validateUserUniqueness(user); err != nil {
		return err
	}

	return s.repo.CreateAdminUser(user)
}

// UpdateAdminUser 更新管理员用户信息
func (s *AdminUserService) UpdateAdminUser(user *models.AdminUser) error {
	if err := s.validateUserUniqueness(user); err != nil {
		return err
	}

	return s.repo.UpdateAdminUser(user)
}

// DeleteAdminUser 删除管理员用户
func (s *AdminUserService) DeleteAdminUser(userID uint) error {
	return s.repo.DeleteAdminUser(userID)
}

// GetAdminUserByID 根据ID获取管理员用户
func (s *AdminUserService) GetAdminUserByID(userID uint) (*models.AdminUser, error) {
	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("用户不存在", err)
		}
		return nil, errors.NewInternalServerError("获取用户失败", err)
	}
	return user, nil
}

// ChangePassword 修改用户密码
func (s *AdminUserService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	user, err := s.GetAdminUserByID(userID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.NewBadRequestError("旧密码不正确", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.NewInternalServerError("密码加密失败", err)
	}

	user.Password = string(hashedPassword)
	return s.repo.UpdateAdminUser(user)
}

func (s *AdminUserService) validateUserUniqueness(user *models.AdminUser) error {
	existingUserByEmail, err := s.repo.FindUserByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.NewInternalServerError("检查邮箱唯一性失败", err)
	}
	if existingUserByEmail != nil && existingUserByEmail.ID != user.ID {
		return errors.NewBadRequestError("邮箱已存在，请使用其他邮箱", nil)
	}

	existingUserByPhone, err := s.repo.FindUserByPhone(user.Phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.NewInternalServerError("检查手机号唯一性失败", err)
	}
	if existingUserByPhone != nil && existingUserByPhone.ID != user.ID {
		return errors.NewBadRequestError("手机号已存在，请使用其他手机号", nil)
	}

	return nil
}

func (s *AdminUserService) Login(email, password string) (string, *models.AdminUser, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		if errors.IsNotFound(err) {
			// 不要明确指出邮箱不存在，而是使用一个通用的错误消息
			return "", nil, errors.NewUnauthorizedError("邮箱或密码错误", nil)
		}
		// 对于其他类型的错误，返回一个通用的服务器错误
		return "", nil, errors.NewInternalServerError("登录过程中发生错误", err)
	}

	if !user.IsActive() {
		// 可以保留这个特定的错误，因为它不会泄露敏感信息
		return "", nil, errors.NewUnauthorizedError("用户账号已被禁用", nil)
	}

	if !user.CheckPassword(password) {
		// 使用与邮箱不存在相同的错误消息
		return "", nil, errors.NewUnauthorizedError("邮箱或密码错误", nil)
	}

	// 生成 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.CustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Minute * 1).Unix(), // 1分钟后过期
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 24小时过期
		},
	})

	// 更新最后登录时间
	now := time.Now()
	user.LastLogin = &now
	if err := s.repo.UpdateAdminUser(user); err != nil {
		return "", nil, errors.NewInternalServerError("更新用户登录时间失败", err)
	}

	tokenString, err := token.SignedString([]byte("123456")) // 使用环境变量存储密钥
	if err != nil {
		return "", nil, errors.NewInternalServerError("生成token失败", err)
	}

	// 清除敏感信息
	user.Password = ""

	return tokenString, user, nil
}

func (s *AdminUserService) GetUserInfo(userID uint) (*models.AdminUser, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, errors.NewNotFoundError("用户不存在", err)
	}
	return user, nil
}
