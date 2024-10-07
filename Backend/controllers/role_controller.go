// controllers/role_controller.go

package controllers

import (
	"net/http"
	"strconv"

	"GOCOM/models"
	"GOCOM/services"
	"GOCOM/utils/errors"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService services.RoleServicer
}

func NewRoleController(roleService services.RoleServicer) *RoleController {
	return &RoleController{roleService}
}

// CreateRole godoc
// @Summary 创建角色
// @Description 创建一个新的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param role body models.Role true "角色信息"
// @Success 201 {object} models.Role "返回创建的角色信息"
// @Failure 400 {object} map[string]string "无效的请求参数"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /api/role [post]
func (c *RoleController) CreateRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.roleService.CreateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, role)
}

// GetAllRoles godoc
// @Summary 获取所有角色
// @Description 获取所有角色，支持分页
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(10)
// @Success 200 {object} map[string]interface{} "返回角色列表、总数、页码和每页数量"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /api/role [get]
func (c *RoleController) GetAllRoles(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	roles, total, err := c.roleService.GetAllRoles(page, pageSize)
	// roles, err := c.roleService.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"roles":    roles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// GetRoleByID godoc
// @Summary 根据ID获取角色
// @Description 通过ID获取特定的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} models.Role "返回角色信息"
// @Failure 400 {object} map[string]string "无效的ID"
// @Failure 404 {object} map[string]string "角色未找到"
// @Router /api/role/{id} [get]
func (c *RoleController) GetRoleByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	role, err := c.roleService.GetRoleByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "角色未找到"})
		return
	}
	ctx.JSON(http.StatusOK, role)
}

// UpdateRole godoc
// @Summary 更新角色
// @Description 更新特定的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param role body models.Role true "更新的角色信息"
// @Success 200 {object} map[string]interface{} "返回更新成功消息和更新后的角色信息"
// @Failure 400 {object} map[string]string "无效的请求参数"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /api/role [put]
func (c *RoleController) UpdateRole(ctx *gin.Context) {
	var role models.Role

	// 将请求中的 JSON 数据绑定到 role 结构体
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.roleService.UpdateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{"message": "更新成功!", "user": role})
}

// DeleteRole godoc
// @Summary 删除角色
// @Description 删除特定的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} map[string]string "删除成功消息"
// @Failure 400 {object} map[string]string "无效的ID"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /api/role/{id} [delete]
func (c *RoleController) DeleteRole(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := c.roleService.DeleteRole(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "角色删除成功"})
}

// AssignMenusToRole godoc
// @Summary 为角色分配菜单
// @Description 为特定角色分配一组菜单
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param request body map[string]interface{} true "包含role_id和menu_ids的JSON"
// @Success 200 {object} map[string]string "菜单分配成功消息"
// @Failure 400 {object} errors.APIError "无效的请求参数"
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/role/assignMenus [post]
func (c *RoleController) AssignMenusToRole(ctx *gin.Context) {
	var req struct {
		RoleID  uint   `json:"role_id" binding:"required"`
		MenuIDs []uint `json:"menu_ids" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的请求参数", err))
		return
	}

	if err := c.roleService.AssignMenusToRole(req.RoleID, req.MenuIDs); err != nil {
		errors.HandleAPIError(ctx, errors.NewInternalServerError("分配菜单失败", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "菜单分配成功"})
}

// GetAssignedMenus godoc
// @Summary 获取角色已分配的菜单
// @Description 获取特定角色已分配的菜单ID列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} map[string][]uint "返回已分配的菜单ID列表"
// @Failure 400 {object} errors.APIError "无效的角色ID"
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/role/{id}/assignedMenus [get]
func (c *RoleController) GetAssignedMenus(ctx *gin.Context) {
	roleID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的角色ID", err))
		return
	}

	menuIDs, err := c.roleService.GetAssignedMenuIDs(uint(roleID))
	if err != nil {
		errors.HandleAPIError(ctx, errors.NewInternalServerError("获取已分配菜单失败", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"menu_ids": menuIDs})
}
