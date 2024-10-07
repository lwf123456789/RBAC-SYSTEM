package models

import "time"

// Role 定义角色模型
type Role struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"size:100;unique;not null" json:"name"`
	Description string       `gorm:"type:text" json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Menus       []Menu       `json:"menus" gorm:"many2many:role_menus;"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
	Users       []*AdminUser `gorm:"many2many:user_roles;joinForeignKey:RoleID;joinReferences:UserID" json:"users"`
}

// TableName 返回数据库表名
func (Role) TableName() string {
	return "roles"
}

// HasPermission 检查角色是否拥有指定权限
func (r *Role) HasPermission(permissionCode string) bool {
	for _, permission := range r.Permissions {
		if permission.Code == permissionCode {
			return true
		}
	}
	return false
}
