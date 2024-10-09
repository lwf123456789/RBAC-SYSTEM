package department

import (
	"GOCOM/config"
	"GOCOM/controllers"
	"GOCOM/middleware"
	"GOCOM/repositories"
	"GOCOM/services"

	"github.com/gin-gonic/gin"
)

func RegisterDepartmentRoutes(r *gin.Engine) {
	departmentRepo := repositories.NewDepartmentRepository(config.DB)
	departmentService := services.NewDepartmentService(departmentRepo)
	departmentController := controllers.NewDepartmentController(departmentService)

	departmentRoutes := r.Group("/api/department")
	departmentRoutes.Use(middleware.JWTAuthMiddleware())
	{
		departmentRoutes.GET("/", departmentController.GetAllDepartments)
		departmentRoutes.POST("/", departmentController.CreateDepartment)
		departmentRoutes.PUT("/:id", departmentController.UpdateDepartment)
		departmentRoutes.DELETE("/:id", departmentController.DeleteDepartment)
	}
}
