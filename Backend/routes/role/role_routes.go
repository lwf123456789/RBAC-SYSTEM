// routes/role_routes.go

package role

import (
	"GOCOM/controllers"
	"GOCOM/middleware"
	"GOCOM/repositories"
	"GOCOM/services"

	"GOCOM/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.Engine) {
	// 创建存储库、服务和控制器
	roleRepo := repositories.NewRoleRepository(config.DB)
	roleService := services.NewRoleService(roleRepo)
	roleController := controllers.NewRoleController(roleService)

	roleRoutes := r.Group("/api/role")
	roleRoutes.Use(middleware.JWTAuthMiddleware())
	{
		roleRoutes.POST("/assignMenus", roleController.AssignMenusToRole)
		roleRoutes.POST("/", roleController.CreateRole)
		roleRoutes.GET("/", roleController.GetAllRoles)
		roleRoutes.GET("/:id", roleController.GetRoleByID)
		roleRoutes.PUT("/", roleController.UpdateRole)
		roleRoutes.DELETE("/:id", roleController.DeleteRole)
		roleRoutes.GET("/:id/assignedMenus", roleController.GetAssignedMenus)
	}
}
