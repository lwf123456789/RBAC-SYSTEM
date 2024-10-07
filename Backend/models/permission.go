package models

import "time"

type Permission struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Code        string    `gorm:"size:100;unique;not null" json:"code"`
	Description string    `gorm:"type:text" json:"description"`
	MenuID      *uint     `json:"menu_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Roles       []Role    `gorm:"many2many:role_permissions;" json:"roles"`
	Menu        *Menu     `gorm:"foreignKey:MenuID" json:"menu"`
	RoleIDs     []uint    `json:"role_ids" gorm:"-"`
}

func (Permission) TableName() string {
	return "permissions"
}
