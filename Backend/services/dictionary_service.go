package services

import (
	"GOCOM/models"
	"GOCOM/repositories"
)

type DictionaryService interface {
	GetDictionariesByType(dictType string) ([]models.Dictionary, error)
	CreateDictionary(dict *models.Dictionary) error
	UpdateDictionary(dict *models.Dictionary) error
	DeleteDictionary(id uint) error
	GetDictionaryByID(id uint) (*models.Dictionary, error)
	GetAllDictionaries(page, pageSize int, typeSearch, labelSearch string) ([]models.Dictionary, int64, error)
}

type dictionaryService struct {
	repo repositories.IDictionaryRepository
}

func NewDictionaryService(repo repositories.IDictionaryRepository) DictionaryService {
	return &dictionaryService{repo: repo}
}

func (ds *dictionaryService) GetDictionariesByType(dictType string) ([]models.Dictionary, error) {
	return ds.repo.GetDictionariesByType(dictType)
}

func (ds *dictionaryService) CreateDictionary(dict *models.Dictionary) error {
	return ds.repo.CreateDictionary(dict)
}

func (ds *dictionaryService) UpdateDictionary(dict *models.Dictionary) error {
	return ds.repo.UpdateDictionary(dict)
}

func (ds *dictionaryService) DeleteDictionary(id uint) error {
	return ds.repo.DeleteDictionary(id)
}

func (ds *dictionaryService) GetDictionaryByID(id uint) (*models.Dictionary, error) {
	return ds.repo.GetDictionaryByID(id)
}

func (ds *dictionaryService) GetAllDictionaries(page, pageSize int, typeSearch, labelSearch string) ([]models.Dictionary, int64, error) {
	return ds.repo.GetAllDictionaries(page, pageSize, typeSearch, labelSearch)
}
