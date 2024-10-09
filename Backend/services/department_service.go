package services

import (
	"GOCOM/models"
	"GOCOM/repositories"
	"GOCOM/utils/errors"
)

type DepartmentService interface {
	GetAllDepartments(page, pageSize int) ([]models.Department, int64, error)
	CreateDepartment(department *models.Department) error
	UpdateDepartment(department *models.Department) (*models.Department, error)
	DeleteDepartment(id uint) error
}

type departmentService struct {
	repo repositories.IDepartmentRepository
}

func NewDepartmentService(repo repositories.IDepartmentRepository) DepartmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) GetAllDepartments(page, pageSize int) ([]models.Department, int64, error) {
	return s.repo.GetAllDepartments(page, pageSize)
}

func (s *departmentService) CreateDepartment(department *models.Department) error {
	return s.repo.CreateDepartment(department)
}

func (s *departmentService) UpdateDepartment(department *models.Department) (*models.Department, error) {
	existingDepartment, err := s.repo.GetDepartmentByID(department.ID)
	if err != nil {
		return nil, err
	}
	if existingDepartment == nil {
		return nil, errors.NewNotFoundError("部门不存在", nil)
	}

	return s.repo.UpdateDepartment(department)
}

func (s *departmentService) DeleteDepartment(id uint) error {
	existingDepartment, err := s.repo.GetDepartmentByID(id)
	if err != nil {
		return err
	}
	if existingDepartment == nil {
		return errors.NewNotFoundError("部门不存在", nil)
	}

	return s.repo.DeleteDepartment(id)
}
