package models

type RoleMenu struct {
	RoleID uint `gorm:"primaryKey"`
	MenuID uint `gorm:"primaryKey"`
}

func (RoleMenu) TableName() string {
	return "role_menus"
}
