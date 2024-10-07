package controllers

import (
	"GOCOM/models"
	"GOCOM/services"
	"GOCOM/utils/errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// AdminUserController 处理管理员用户相关的HTTP请求
type AdminUserController struct {
	adminUserService services.AdminUserServicer
}

// NewAdminUserController 创建一个新的AdminUserController实例
func NewAdminUserController(adminUserService services.AdminUserServicer) *AdminUserController {
	return &AdminUserController{
		adminUserService: adminUserService,
	}
}

// GetAllUsersWithRoles godoc
// @Summary 获取用户列表及其角色
// @Description 获取所有用户及其角色信息，支持分页、按部门过滤、邮箱和名字模糊搜索，以及按角色筛选
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param page query int false "当前页码" default(1)
// @Param pageSize query int false "每页记录数量" default(10)
// @Param departmentID query int false "部门ID，用于过滤该部门下的用户"
// @Param email query string false "邮箱模糊搜索"
// @Param name query string false "名字模糊搜索"
// @Param roles query string false "角色ID列表，用逗号分隔，如 '3,4'"
// @Success 200 {object} map[string]interface{} "返回包含用户列表、总记录数、当前页和每页数量的 JSON 数据"
// @Failure 400 {object} errors.APIError "请求参数错误"
// @Failure 500 {object} errors.APIError "获取用户数据失败"
// @Router /api/adminUser/get [get]
func (c *AdminUserController) GetAllUsersWithRoles(ctx *gin.Context) {
	params, err := c.parseGetAllUsersParams(ctx)
	if err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的请求参数", err))
		return
	}

	users, total, err := c.adminUserService.GetAllUsersWithRoles(
		params.Page,
		params.PageSize,
		params.DepartmentID,
		params.Email,
		params.Name,
		params.RoleIDs,
	)
	if err != nil {
		errors.HandleAPIError(ctx, errors.NewInternalServerError("获取用户数据失败", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users":    users,
		"total":    total,
		"page":     params.Page,
		"pageSize": params.PageSize,
	})
}

type getAllUsersParams struct {
	Page         int
	PageSize     int
	DepartmentID *uint
	Email        string
	Name         string
	RoleIDs      []uint
}

func (c *AdminUserController) parseGetAllUsersParams(ctx *gin.Context) (getAllUsersParams, error) {
	params := getAllUsersParams{
		Page:     1,
		PageSize: 10,
	}

	if page, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err == nil && page > 0 {
		params.Page = page
	}

	if pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10")); err == nil && pageSize > 0 {
		params.PageSize = pageSize
	}

	if departmentIDStr := ctx.Query("departmentID"); departmentIDStr != "" {
		if departmentID, err := strconv.ParseUint(departmentIDStr, 10, 32); err == nil {
			deptID := uint(departmentID)
			params.DepartmentID = &deptID
		}
	}

	params.Email = ctx.Query("email")
	params.Name = ctx.Query("name")

	// 处理角色 ID
	rolesStr := ctx.Query("roles")
	if rolesStr != "" {
		roleIDs := []uint{}
		for _, idStr := range strings.Split(rolesStr, ",") {
			if id, err := strconv.ParseUint(idStr, 10, 32); err == nil {
				roleIDs = append(roleIDs, uint(id))
			}
		}
		params.RoleIDs = roleIDs
	}

	return params, nil
}

// CreateAdminUser godoc
// @Summary 创建新用户
// @Description 创建一个新的用户，并关联角色
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param user body models.AdminUser true "用户信息"
// @Success 200 {object} map[string]interface{} "返回创建的用户信息"
// @Failure 400 {object} errors.APIError "创建用户失败"
// @Router /api/admin_user [post]
func (c *AdminUserController) CreateAdminUser(ctx *gin.Context) {
	var adminUser models.AdminUser
	if err := ctx.ShouldBindJSON(&adminUser); err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的用户数据", err))
		return
	}

	if err := c.adminUserService.CreateAdminUser(&adminUser); err != nil {
		errors.HandleAPIError(ctx, errors.NewInternalServerError("创建用户失败", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "创建成功", "user": adminUser})
}

// UpdateAdminUser godoc
// @Summary 更新用户信息
// @Description 更新指定用户的信息
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body models.AdminUser true "用户信息"
// @Success 200 {object} map[string]interface{} "返回更新后的用户信息"
// @Failure 400 {object} errors.APIError "请求参数错误"
// @Failure 500 {object} errors.APIError "服务器错误"
// @Router /api/admin_user/{id} [put]
func (c *AdminUserController) UpdateAdminUser(ctx *gin.Context) {
	var adminUser models.AdminUser
	if err := ctx.ShouldBindJSON(&adminUser); err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的用户数据", err))
		return
	}

	if err := c.adminUserService.UpdateAdminUser(&adminUser); err != nil {
		errors.HandleAPIError(ctx, errors.NewInternalServerError("更新用户失败", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "更新成功", "user": adminUser})
}

// DeleteAdminUser godoc
// @Summary 删除用户
// @Description 删除指定用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} map[string]interface{} "返回删除成功信息"
// @Failure 400 {object} errors.APIError "请求参数错误"
// @Failure 500 {object} errors.APIError "服务器错误"
// @Router /api/admin_user/{id} [delete]
func (c *AdminUserController) DeleteAdminUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的用户ID", err))
		return
	}

	if err := c.adminUserService.DeleteAdminUser(uint(id)); err != nil {
		errors.HandleAPIError(ctx, errors.NewInternalServerError("删除用户失败", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login godoc
// @Summary 用户登录
// @Description 管理员用户登录并获取token
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "登录信息"
// @Success 200 {object} map[string]interface{} "返回token和用户信息"
// @Failure 400 {object} errors.APIError "无效的请求参数"
// @Failure 401 {object} errors.APIError "登录失败"
// @Router /api/admin_user/login [post]
func (c *AdminUserController) Login(ctx *gin.Context) {
	var loginReq LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		errors.HandleAPIError(ctx, errors.NewBadRequestError("无效的请求参数", err))
		return
	}

	token, user, err := c.adminUserService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		errors.HandleAPIError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

// GetUserInfo godoc
// @Summary 获取用户信息
// @Description 获取当前登录用户的详细信息
// @Tags 用户模块
// @Accept json
// @Produce json
// @Success 200 {object} models.AdminUser "返回用户信息"
// @Failure 401 {object} errors.APIError "用户未认证"
// @Failure 500 {object} errors.APIError "获取用户信息失败"
// @Router /api/admin_user/info [get]
func (c *AdminUserController) GetUserInfo(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")

	if !exists {
		errors.HandleAPIError(ctx, errors.NewUnauthorizedError("用户未认证", nil))
		return
	}

	user, err := c.adminUserService.GetUserInfo(userID.(uint))
	if err != nil {
		errors.HandleAPIError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
