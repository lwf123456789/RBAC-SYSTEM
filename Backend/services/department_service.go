package services

import (
	"GOCOM/models"
	"GOCOM/repositories"
)

type DepartmentService interface {
	GetAllDepartments() ([]models.Department, error)
}

type departmentService struct {
	repo repositories.IDepartmentRepository
}

func NewDepartmentService(repo repositories.IDepartmentRepository) DepartmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) GetAllDepartments() ([]models.Department, error) {
	return s.repo.GetAllDepartments()
}
