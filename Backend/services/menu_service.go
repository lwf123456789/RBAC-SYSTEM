package services

import (
	"GOCOM/models"
	"GOCOM/repositories"
)

type MenuServicer interface {
	CreateMenu(menu *models.Menu) error
	GetAllMenus(page, pageSize int) ([]models.Menu, int64, error)
	GetAllMenusWithRoles(page, pageSize int) ([]models.Menu, int64, error)
	GetMenuByID(id uint) (*models.Menu, error)
	UpdateMenu(menu *models.Menu) error
	DeleteMenu(id uint) error
	GetUserMenus(userID uint) ([]models.Menu, error)
	AddMenuPermissions(menuID uint, permissions []models.Permission) error
	GetMenuPermissions(menuID uint) ([]models.Permission, error)
}

type MenuService struct {
	repo repositories.IMenuRepository
}

func NewMenuService(repo repositories.IMenuRepository) MenuServicer {
	return &MenuService{repo: repo}
}

func (s *MenuService) CreateMenu(menu *models.Menu) error {
	return s.repo.Create(menu)
}

func (s *MenuService) GetAllMenus(page, pageSize int) ([]models.Menu, int64, error) {
	return s.repo.GetAll(page, pageSize)
}

func (s *MenuService) GetAllMenusWithRoles(page, pageSize int) ([]models.Menu, int64, error) {
	return s.repo.GetAllWithRoles(page, pageSize)
}

func (s *MenuService) GetMenuByID(id uint) (*models.Menu, error) {
	return s.repo.GetByID(id)
}

func (s *MenuService) UpdateMenu(menu *models.Menu) error {
	return s.repo.Update(menu)
}

func (s *MenuService) DeleteMenu(id uint) error {
	return s.repo.Delete(id)
}

func (s *MenuService) GetUserMenus(userID uint) ([]models.Menu, error) {
	return s.repo.GetUserMenus(userID)
}

func (s *MenuService) AddMenuPermissions(menuID uint, permissions []models.Permission) error {
	return s.repo.AddMenuPermissions(menuID, permissions)
}

func (s *MenuService) GetMenuPermissions(menuID uint) ([]models.Permission, error) {
	return s.repo.GetMenuPermissions(menuID)
}
