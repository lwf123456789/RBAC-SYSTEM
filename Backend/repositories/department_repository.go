package repositories

import (
	"GOCOM/models"
	"GOCOM/utils/errors"

	"gorm.io/gorm"
)

type IDepartmentRepository interface {
	GetAllDepartments(page, pageSize int) ([]models.Department, int64, error)
	CreateDepartment(department *models.Department) error
	UpdateDepartment(department *models.Department) (*models.Department, error)
	DeleteDepartment(id uint) error
	GetDepartmentByID(id uint) (*models.Department, error)
}

type DepartmentRepository struct {
	DB *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) IDepartmentRepository {
	return &DepartmentRepository{DB: db}
}

func (r *DepartmentRepository) GetAllDepartments(page, pageSize int) ([]models.Department, int64, error) {
	var departments []models.Department
	var total int64

	offset := (page - 1) * pageSize

	err := r.DB.Model(&models.Department{}).Count(&total).Error
	if err != nil {
		return nil, 0, errors.NewInternalServerError("获取部门总数失败", err)
	}

	err = r.DB.Offset(offset).Limit(pageSize).Find(&departments).Error
	if err != nil {
		return nil, 0, errors.NewInternalServerError("获取部门数据失败", err)
	}

	return departments, total, nil
}

func (r *DepartmentRepository) CreateDepartment(department *models.Department) error {
	err := r.DB.Create(department).Error
	if err != nil {
		return errors.NewInternalServerError("创建部门失败", err)
	}
	return nil
}

func (r *DepartmentRepository) UpdateDepartment(department *models.Department) (*models.Department, error) {
	err := r.DB.Omit("created_at").Updates(department).Error
	if err != nil {
		return nil, errors.NewInternalServerError("更新部门失败", err)
	}
	return department, nil
}

func (r *DepartmentRepository) DeleteDepartment(id uint) error {
	err := r.DB.Delete(&models.Department{}, id).Error
	if err != nil {
		return errors.NewInternalServerError("删除部门失败", err)
	}
	return nil
}

func (r *DepartmentRepository) GetDepartmentByID(id uint) (*models.Department, error) {
	var department models.Department
	err := r.DB.First(&department, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.NewInternalServerError("获取部门失败", err)
	}
	return &department, nil
}
