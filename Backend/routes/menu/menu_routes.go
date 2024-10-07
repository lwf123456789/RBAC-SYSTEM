package menu

import (
	"GOCOM/config"
	"GOCOM/controllers"
	"GOCOM/middleware"
	"GOCOM/repositories"
	"GOCOM/services"

	"github.com/gin-gonic/gin"
)

func RegisterMenuRoutes(r *gin.Engine) {
	menuRepo := repositories.NewMenuRepository(config.DB)
	menuService := services.NewMenuService(menuRepo)
	menuController := controllers.NewMenuController(menuService)

	menuGroup := r.Group("/api/menus")
	menuGroup.Use(middleware.JWTAuthMiddleware())
	{
		menuGroup.GET("/getRoutes", menuController.GetUserMenus)
		menuGroup.POST("/", menuController.CreateMenu)
		menuGroup.GET("/", menuController.GetAllMenus)
		menuGroup.GET("/menusWithRoles", menuController.GetAllMenusWithRoles)
		menuGroup.GET("/:id", menuController.GetMenuByID)
		menuGroup.PUT("/", menuController.UpdateMenu)
		menuGroup.DELETE("/:id", menuController.DeleteMenu)
		menuGroup.POST("/permissions/:id", menuController.AddMenuPermissions)
		menuGroup.GET("/permissions/:id", menuController.GetMenuPermissions)
	}
}
