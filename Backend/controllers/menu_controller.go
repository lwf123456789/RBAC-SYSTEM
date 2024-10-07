package controllers

import (
	"GOCOM/models"
	"GOCOM/services"
	"GOCOM/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	service services.MenuServicer
}

func NewMenuController(service services.MenuServicer) *MenuController {
	return &MenuController{service: service}
}

// CreateMenu godoc
// @Summary 创建新菜单
// @Description 创建一个新的菜单项
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param menu body models.Menu true "菜单信息"
// @Success 201 {object} models.Menu "返回创建的菜单信息"
// @Failure 400 {object} map[string]string "无效的请求参数"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /api/menus [post]
func (c *MenuController) CreateMenu(ctx *gin.Context) {
	var menu models.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateMenu(&menu); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, menu)
}

// GetAllMenus godoc
// @Summary 获取所有菜单
// @Description 获取所有菜单项，支持分页
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(10)
// @Success 200 {object} map[string]interface{} "返回菜单列表、总数、页码和每页数量"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /api/menus [get]
func (c *MenuController) GetAllMenus(ctx *gin.Context) {
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

	menus, total, err := c.service.GetAllMenus(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"menus":    menus,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// GetAllMenusWithRoles godoc
// @Summary 获取所有菜单及其角色
// @Description 获取所有菜单项及其关联的角色，支持分页
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(10)
// @Success 200 {object} map[string]interface{} "返回菜单列表（包含角色）、总数、页码和每页数量"
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/menus/menusWithRoles [get]
func (c *MenuController) GetAllMenusWithRoles(ctx *gin.Context) {
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

	menus, total, err := c.service.GetAllMenusWithRoles(page, pageSize)
	if err != nil {
		errors.HandleAPIError(ctx, errors.NewInternalServerError("获取菜单列表失败", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"menus":    menus,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// GetMenuByID godoc
// @Summary 通过ID获取菜单
// @Description 通过ID获取特定的菜单项
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Success 200 {object} models.Menu "返回菜单信息"
// @Failure 400 {object} map[string]string "无效的ID"
// @Failure 404 {object} map[string]string "菜单未找到"
// @Router /api/menus/{id} [get]
func (c *MenuController) GetMenuByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	menu, err := c.service.GetMenuByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

// UpdateMenu godoc
// @Summary 更新菜单
// @Description 更新特定的菜单项
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param menu body models.Menu true "更新的菜单信息"
// @Success 200 {object} models.Menu "返回更新后的菜单信息"
// @Failure 400 {object} map[string]string "无效的请求参数"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /api/menus [put]
func (c *MenuController) UpdateMenu(ctx *gin.Context) {
	var menu models.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateMenu(&menu); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

// DeleteMenu godoc
// @Summary 删除菜单
// @Description 删除特定的菜单项
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Success 200 {object} map[string]string "删除成功消息"
// @Failure 400 {object} map[string]string "无效的ID"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /api/menus/{id} [delete]
func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.service.DeleteMenu(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}

// GetUserMenus godoc
// @Summary 获取用户菜单
// @Description 获取当前登录用户的菜单列表
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Success 200 {array} models.Menu "返回用户的菜单列表"
// @Failure 401 {object} errors.APIError "用户未认证"
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/menus/getRoutes [get]
func (c *MenuController) GetUserMenus(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		errors.HandleAPIError(ctx, errors.NewUnauthorizedError("用户未认证", nil))
		return
	}

	menus, err := c.service.GetUserMenus(userID.(uint))
	if err != nil {
		errors.HandleAPIError(ctx, errors.NewInternalServerError("获取用户菜单失败", err))
		return
	}

	ctx.JSON(http.StatusOK, menus)
}

// AddMenuPermissions godoc
// @Summary 为菜单添加权限
// @Description 为指定菜单添加一组权限
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Param permissions body []models.Permission true "权限列表"
// @Success 200 {object} map[string]string "添加权限成功"
// @Failure 400 {object} errors.APIError "无效的请求参数"
// @Failure 404 {object} errors.APIError "菜单未找到"
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/menus/permissions/{id} [post]
func (c *MenuController) AddMenuPermissions(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的菜单ID", err))
		return
	}

	var requestBody struct {
		Permissions []models.Permission `json:"permissions"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的请求参数", err))
		return
	}

	err = c.service.AddMenuPermissions(uint(id), requestBody.Permissions)
	if err != nil {
		if errors.IsNotFound(err) {
			errors.HandleAPIError(ctx, errors.NewNotFoundError("菜单未找到", err))
		} else {
			errors.HandleAPIError(ctx, errors.NewInternalServerError("添加权限失败", err))
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "权限添加成功"})
}

// GetMenuPermissions godoc
// @Summary 获取指定菜单的所有权限
// @Description 获取指定菜单ID下的所有权限数据
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param id path int true "菜单ID"
// @Success 200 {array} models.Permission "返回权限列表"
// @Failure 400 {object} errors.APIError "无效的菜单ID"
// @Failure 404 {object} errors.APIError "菜单未找到"
// @Failure 500 {object} errors.APIError "服务器内部错误"
// @Router /api/menus/{id}/permissions [get]
func (c *MenuController) GetMenuPermissions(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的菜单ID", err))
		return
	}

	permissions, err := c.service.GetMenuPermissions(uint(id))
	if err != nil {
		if errors.IsNotFound(err) {
			errors.HandleAPIError(ctx, errors.NewNotFoundError("菜单未找到", err))
		} else {
			errors.HandleAPIError(ctx, errors.NewInternalServerError("获取菜单权限失败", err))
		}
		return
	}

	ctx.JSON(http.StatusOK, permissions)
}
