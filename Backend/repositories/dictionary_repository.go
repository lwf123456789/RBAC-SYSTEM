package repositories

import (
	"GOCOM/models"

	"gorm.io/gorm"
)

type IDictionaryRepository interface {
	GetDictionariesByType(dictType string) ([]models.Dictionary, error)
	CreateDictionary(dict *models.Dictionary) error
	UpdateDictionary(dict *models.Dictionary) error
	DeleteDictionary(id uint) error
	GetDictionaryByID(id uint) (*models.Dictionary, error)
	GetAllDictionaries(page, pageSize int, typeSearch, labelSearch string) ([]models.Dictionary, int64, error)
}

type DictionaryRepository struct {
	DB *gorm.DB
}

func NewDictionaryRepository(db *gorm.DB) IDictionaryRepository {
	return &DictionaryRepository{DB: db}
}

func (repo *DictionaryRepository) GetDictionariesByType(dictType string) ([]models.Dictionary, error) {
	var dictionaries []models.Dictionary
	if err := repo.DB.Where("type = ?", dictType).
		Order("sort ASC, id DESC").
		Find(&dictionaries).Error; err != nil {
		return nil, err
	}
	return dictionaries, nil
}

func (repo *DictionaryRepository) CreateDictionary(dict *models.Dictionary) error {
	return repo.DB.Create(dict).Error
}

func (repo *DictionaryRepository) UpdateDictionary(dict *models.Dictionary) error {
	return repo.DB.Omit("created_at").Updates(dict).Error
}

func (repo *DictionaryRepository) DeleteDictionary(id uint) error {
	return repo.DB.Delete(&models.Dictionary{}, id).Error
}

func (repo *DictionaryRepository) GetDictionaryByID(id uint) (*models.Dictionary, error) {
	var dict models.Dictionary
	if err := repo.DB.First(&dict, id).Error; err != nil {
		return nil, err
	}
	return &dict, nil
}

func (repo *DictionaryRepository) GetAllDictionaries(page, pageSize int, typeSearch, labelSearch string) ([]models.Dictionary, int64, error) {
	var dictionaries []models.Dictionary
	var total int64

	offset := (page - 1) * pageSize

	query := repo.DB.Model(&models.Dictionary{})

	if typeSearch != "" {
		query = query.Where("type LIKE ?", "%"+typeSearch+"%")
	}
	if labelSearch != "" {
		query = query.Where("label LIKE ?", "%"+labelSearch+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&dictionaries).Error; err != nil {
		return nil, 0, err
	}

	return dictionaries, total, nil
}
