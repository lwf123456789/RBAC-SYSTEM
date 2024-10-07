// services/role_service.go

package services

import (
	"GOCOM/models"
	"GOCOM/repositories"
)

type RoleServicer interface {
	CreateRole(role *models.Role) error
	GetAllRoles(page, pageSize int) ([]models.Role, int64, error)
	GetRoleByID(id uint) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
	AssignMenusToRole(roleID uint, menuIDs []uint) error
	GetAssignedMenuIDs(roleID uint) ([]uint, error)
}

type RoleService struct {
	roleRepo repositories.IRoleRepository
}

func NewRoleService(roleRepo repositories.IRoleRepository) RoleServicer {
	return &RoleService{roleRepo: roleRepo}
}

func (s *RoleService) CreateRole(role *models.Role) error {
	return s.roleRepo.CreateRole(role)
}

func (s *RoleService) GetAllRoles(page, pageSize int) ([]models.Role, int64, error) {
	return s.roleRepo.GetAllRoles(page, pageSize)
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	return s.roleRepo.GetRoleByID(id)
}

func (s *RoleService) UpdateRole(role *models.Role) error {
	return s.roleRepo.UpdateRole(role)
}

func (s *RoleService) DeleteRole(id uint) error {
	return s.roleRepo.DeleteRole(id)
}

func (s *RoleService) AssignMenusToRole(roleID uint, menuIDs []uint) error {
	return s.roleRepo.AssignMenusToRole(roleID, menuIDs)
}

func (s *RoleService) GetAssignedMenuIDs(roleID uint) ([]uint, error) {
	return s.roleRepo.GetAssignedMenuIDs(roleID)
}
