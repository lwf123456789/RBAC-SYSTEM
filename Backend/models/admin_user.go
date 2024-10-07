package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminUser struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Name         string     `gorm:"size:255;not null" json:"name"`
	Email        string     `gorm:"size:255;unique;not null" json:"email"`
	Password     string     `gorm:"size:255;not null" json:"-"` // 使用 json:"-" 确保密码不会被序列化
	Phone        string     `gorm:"size:20" json:"phone"`
	DepartmentID *uint      `json:"department_id"`
	Status       int        `gorm:"type:tinyint;default:1" json:"status"`
	LastLogin    *time.Time `json:"last_login"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	Roles      []*Role     `gorm:"many2many:user_roles;joinForeignKey:UserID;joinReferences:RoleID" json:"roles"`
	Department *Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
}

// TableName 返回数据库表名
func (AdminUser) TableName() string {
	return "admin_users"
}

// BeforeCreate GORM 钩子，在创建用户前加密密码
func (u *AdminUser) BeforeCreate(tx *gorm.DB) error {
	return u.hashPassword()
}

// BeforeUpdate GORM 钩子，在更新用户前加密密码（如果密码被更改）
func (u *AdminUser) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("Password") {
		return u.hashPassword()
	}
	return nil
}

// hashPassword 加密用户密码
func (u *AdminUser) hashPassword() error {
	if len(u.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword 直接比较明文密码
func (u *AdminUser) CheckPassword(password string) bool {
	return u.Password == password
}

// IsActive 检查用户是否处于活动状态
func (u *AdminUser) IsActive() bool {
	return u.Status == 1
}

// HasRole 检查用户是否拥有指定角色
func (u *AdminUser) HasRole(roleName string) bool {
	for _, role := range u.Roles {
		if role.Name == roleName {
			return true
		}
	}
	return false
}

// HasPermission 检查用户是否拥有指定权限
func (u *AdminUser) HasPermission(permissionCode string) bool {
	for _, role := range u.Roles {
		if role.HasPermission(permissionCode) {
			return true
		}
	}
	return false
}

// IsSuperAdmin 检查用户是否为超级管理员
func (u *AdminUser) IsSuperAdmin() bool {
	for _, role := range u.Roles {
		if role.ID == 1 { // 假设超级管理员角色的 ID 为 1
			return true
		}
	}
	return false
}
