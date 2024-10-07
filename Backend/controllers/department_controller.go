package controllers

import (
	"GOCOM/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	service services.DepartmentService
}

func NewDepartmentController(service services.DepartmentService) *DepartmentController {
	return &DepartmentController{service: service}
}

// GetCarousels godoc
// @Summary 获取部门列表
// @Description 获取部门所有数据
// @Tags 部门管理
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Department
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /api/department [get]
func (ctrl *DepartmentController) GetAllDepartments(c *gin.Context) {
	departments, err := ctrl.service.GetAllDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取部门数据"})
		return
	}
	c.JSON(http.StatusOK, departments)
}
