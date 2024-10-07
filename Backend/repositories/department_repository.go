package repositories

import (
	"GOCOM/models"
	"GOCOM/utils/errors"

	"gorm.io/gorm"
)

type IDepartmentRepository interface {
	GetAllDepartments() ([]models.Department, error)
}

type DepartmentRepository struct {
	DB *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) IDepartmentRepository {
	return &DepartmentRepository{DB: db}
}

func (r *DepartmentRepository) GetAllDepartments() ([]models.Department, error) {
	var departments []models.Department
	err := r.DB.Find(&departments).Error
	if err != nil {
		return nil, errors.NewInternalServerError("获取部门数据失败", err)
	}

	return departments, nil
}
