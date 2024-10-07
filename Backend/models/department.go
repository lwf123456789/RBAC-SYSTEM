package models

import (
	"time"
)

type Department struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	Name      string       `gorm:"size:255;not null" json:"name"`
	ParentID  *uint        `json:"parent_id"`
	Icon      *string      `gorm:"size:255" json:"icon"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Children  []Department `gorm:"-" json:"children,omitempty"`
}

// TableName 指定表名
func (Department) TableName() string {
	return "departments"
}
