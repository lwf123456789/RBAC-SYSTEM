package controllers

import (
	"GOCOM/models"
	"GOCOM/services"
	"GOCOM/utils/errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DictionaryController struct {
	service services.DictionaryService
}

func NewDictionaryController(service services.DictionaryService) *DictionaryController {
	return &DictionaryController{service: service}
}

// GetDictionaries godoc
// @Summary 获取数据字典
// @Description 根据类型获取数据字典
// @Tags 字典管理
// @Accept  json
// @Produce  json
// @Param   type   query    string  true  "字典类型"
// @Success 200 {array} models.Dictionary
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Router /api/dicts [get]
func (dc *DictionaryController) GetDictionaries(c *gin.Context) {
	dictType := c.Query("type")
	if dictType == "" {
		errors.HandleAPIError(c, errors.NewBadRequestError("Type is required", nil))
		return
	}

	dictionaries, err := dc.service.GetDictionariesByType(dictType)
	if err != nil {
		errors.HandleAPIError(c, errors.NewInternalServerError("Failed to fetch dictionaries", err))
		return
	}

	c.JSON(http.StatusOK, dictionaries)
}

// CreateDictionary godoc
// @Summary 创建数据字典
// @Description 创建新的数据字典项
// @Tags 字典管理
// @Accept  json
// @Produce  json
// @Param   dictionary body models.Dictionary true "数据字典信息"
// @Success 201 {object} models.Dictionary
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Router /api/dicts [post]
func (dc *DictionaryController) CreateDictionary(c *gin.Context) {
	var dict models.Dictionary
	if err := c.ShouldBindJSON(&dict); err != nil {
		errors.HandleAPIError(c, errors.NewBadRequestError("Invalid input", err))
		return
	}

	if err := dc.service.CreateDictionary(&dict); err != nil {
		errors.HandleAPIError(c, errors.NewInternalServerError("Failed to create dictionary", err))
		return
	}

	c.JSON(http.StatusCreated, dict)
}

// UpdateDictionary godoc
// @Summary 更新数据字典
// @Description 更新现有的数据字典项
// @Tags 字典管理
// @Accept  json
// @Produce  json
// @Param   id path int true "字典ID"
// @Param   dictionary body models.Dictionary true "数据字典信息"
// @Success 200 {object} models.Dictionary
// @Failure 400 {object} errors.APIError
// @Failure 404 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Router /api/dicts [put]
func (dc *DictionaryController) UpdateDictionary(c *gin.Context) {
	var dict models.Dictionary
	if err := c.ShouldBindJSON(&dict); err != nil {
		fmt.Printf("JSON binding error: %v\n", err)
		errors.HandleAPIError(c, errors.NewBadRequestError("无效的输入", err))
		return
	}

	// 打印接收到的数据
	fmt.Printf("接收到的数据: %+v\n", dict)

	if dict.ID == 0 {
		errors.HandleAPIError(c, errors.NewBadRequestError("ID 是必需的", nil))
		return
	}

	if err := dc.service.UpdateDictionary(&dict); err != nil {
		fmt.Printf("更新错误: %v\n", err)
		errors.HandleAPIError(c, errors.NewInternalServerError("更新字典失败", err))
		return
	}

	c.JSON(http.StatusOK, dict)
}

// DeleteDictionary godoc
// @Summary 删除数据字典
// @Description 删除指定的数据字典项
// @Tags 字典管理
// @Accept  json
// @Produce  json
// @Param   id path int true "字典ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} errors.APIError
// @Failure 500 {object} errors.APIError
// @Router /api/dicts/{id} [delete]
func (dc *DictionaryController) DeleteDictionary(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := dc.service.DeleteDictionary(uint(id)); err != nil {
		errors.HandleAPIError(c, errors.NewInternalServerError("Failed to delete dictionary", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// GetAllDictionaries godoc
// @Summary 获取所有数据字典
// @Description 获取所有数据字典项，支持分页
// @Tags 字典管理
// @Accept  json
// @Produce  json
// @Param   page query int false "页码" default(1)
// @Param   pageSize query int false "每页数量" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} errors.APIError
// @Router /api/dicts/all [get]
func (dc *DictionaryController) GetAllDictionaries(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	typeSearch := c.Query("type")
	labelSearch := c.Query("label")

	dictionaries, total, err := dc.service.GetAllDictionaries(page, pageSize, typeSearch, labelSearch)
	if err != nil {
		errors.HandleAPIError(c, errors.NewInternalServerError("获取字典列表失败", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"dictionaries": dictionaries,
		"total":        total,
		"page":         page,
		"pageSize":     pageSize,
	})
}
