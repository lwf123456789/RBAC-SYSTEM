package admin

import (
	"GOCOM/config"
	"GOCOM/controllers"
	"GOCOM/middleware"
	"GOCOM/repositories"
	"GOCOM/services"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.Engine) {
	// 创建存储库、服务和控制器
	adminUserRepo := repositories.NewAdminUserRepository(config.DB)
	adminUserService := services.NewAdminUserService(adminUserRepo)
	adminUserController := controllers.NewAdminUserController(adminUserService)

	// 公开路由，不需要认证
	r.POST("/api/adminUser/login", adminUserController.Login)

	// 需要认证的路由
	adminUserRoutes := r.Group("/api/adminUser")
	adminUserRoutes.Use(middleware.JWTAuthMiddleware())
	{
		adminUserRoutes.GET("/", adminUserController.GetAllUsersWithRoles)
		adminUserRoutes.POST("/", adminUserController.CreateAdminUser)
		adminUserRoutes.PUT("/", adminUserController.UpdateAdminUser)
		adminUserRoutes.DELETE("/:id", adminUserController.DeleteAdminUser)
		adminUserRoutes.GET("/info", adminUserController.GetUserInfo)
	}
}
