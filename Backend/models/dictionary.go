package models

import (
	"time"
)

type Dictionary struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Type        string    `gorm:"type:varchar(50);not null;comment:字典类型" json:"type"`
	Label       string    `gorm:"type:varchar(100);not null;comment:显示名称" json:"label"`
	Value       string    `gorm:"type:varchar(100);not null;comment:值" json:"value"`
	Sort        int       `gorm:"type:int;default:0;comment:排序" json:"sort"`
	Description string    `gorm:"type:text;comment:描述" json:"description"`
	CreatedAt   time.Time `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time `gorm:"comment:更新时间" json:"updated_at"`
}

// TableName 返回数据库表名
func (Dictionary) TableName() string {
	return "dictionaries"
}
