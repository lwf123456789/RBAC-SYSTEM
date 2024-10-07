package models

import (
	"encoding/json"
	"time"
)

type Menu struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Title         string    `gorm:"size:100;not null" json:"title"`
	Path          string    `gorm:"size:255" json:"path"`
	ComponentPath string    `gorm:"size:255" json:"component_path"`
	Icon          string    `gorm:"size:255" json:"icon"`
	ParentID      *uint     `json:"parent_id"`
	Sort          int       `json:"sort"`
	Status        int       `gorm:"type:tinyint;default:1" json:"status"`
	Permissions   *string   `json:"permissions" gorm:"type:json"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Children      []Menu    `json:"children" gorm:"-"`

	Roles           []*Role      `gorm:"many2many:role_menus;joinForeignKey:MenuID;joinReferences:RoleID" json:"roles,omitempty"`
	PermissionItems []Permission `json:"permission_items" gorm:"foreignKey:MenuID"`
}

// TableName 返回数据库表名
func (Menu) TableName() string {
	return "menus"
}

// SetPermissions 设置权限
func (m *Menu) SetPermissions(permissions []string) error {
	if len(permissions) == 0 {
		m.Permissions = nil
		return nil
	}
	jsonData, err := json.Marshal(permissions)
	if err != nil {
		return err
	}
	strData := string(jsonData)
	m.Permissions = &strData
	return nil
}

// GetPermissions 获取权限
func (m *Menu) GetPermissions() ([]string, error) {
	if m.Permissions == nil {
		return []string{}, nil
	}
	var permissions []string
	err := json.Unmarshal([]byte(*m.Permissions), &permissions)
	return permissions, err
}
