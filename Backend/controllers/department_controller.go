package controllers

import (
	"GOCOM/models"
	"GOCOM/services"
	"GOCOM/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	service services.DepartmentService
}

func NewDepartmentController(service services.DepartmentService) *DepartmentController {
	return &DepartmentController{service: service}
}

// GetAllDepartments godoc
// @Summary 获取所有部门
// @Description 获取所有部门数据
// @Tags 部门管理
// @Accept json
// @Produce json
// @Success 200 {array} models.Department
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/department [get]
func (ctrl *DepartmentController) GetAllDepartments(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "999")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	departments, total, err := ctrl.service.GetAllDepartments(page, pageSize)
	if err != nil {
		errors.HandleAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"departments": departments,
		"total":       total,
		"page":        page,
		"pageSize":    pageSize,
	})
}

// CreateDepartment godoc
// @Summary 创建部门
// @Description 创建一个新的部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param department body models.Department true "部门信息"
// @Success 201 {object} models.Department
// @Failure 400 {object} errors.APIError "无效的请求参数"
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/department [post]
func (ctrl *DepartmentController) CreateDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		errors.HandleAPIError(c, errors.NewBadRequestError("无效的部门数据", err))
		return
	}

	err := ctrl.service.CreateDepartment(&department)
	if err != nil {
		errors.HandleAPIError(c, err)
		return
	}

	c.JSON(http.StatusCreated, department)
}

// UpdateDepartment godoc
// @Summary 更新部门
// @Description 更新指定ID的部门信息
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门ID"
// @Param department body models.Department true "部门信息"
// @Success 200 {object} models.Department
// @Failure 400 {object} errors.APIError "无效的请求参数"
// @Failure 404 {object} errors.APIError "部门不存在"
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/department/{id} [put]
func (ctrl *DepartmentController) UpdateDepartment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errors.HandleAPIError(c, errors.NewBadRequestError("无效的部门ID", err))
		return
	}

	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		errors.HandleAPIError(c, errors.NewBadRequestError("无效的部门数据", err))
		return
	}
	department.ID = uint(id)

	updatedDepartment, err := ctrl.service.UpdateDepartment(&department)
	if err != nil {
		errors.HandleAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedDepartment)
}

// DeleteDepartment godoc
// @Summary 删除部门
// @Description 删除指定ID的部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门ID"
// @Success 204 "No Content"
// @Failure 400 {object} errors.APIError "无效的请求参数"
// @Failure 404 {object} errors.APIError "部门不存在"
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/department/{id} [delete]
func (ctrl *DepartmentController) DeleteDepartment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errors.HandleAPIError(c, errors.NewBadRequestError("无效的部门ID", err))
		return
	}

	err = ctrl.service.DeleteDepartment(uint(id))
	if err != nil {
		errors.HandleAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
