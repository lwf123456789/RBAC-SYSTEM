package dictionary

import (
	"GOCOM/config"
	"GOCOM/controllers"
	"GOCOM/middleware"
	"GOCOM/repositories"
	"GOCOM/services"

	"github.com/gin-gonic/gin"
)

func RegisterDictRoutes(r *gin.Engine) {
	dictRepository := repositories.NewDictionaryRepository(config.DB)
	dictService := services.NewDictionaryService(dictRepository)
	dictController := controllers.NewDictionaryController(dictService)

	dictRoutes := r.Group("/api/dicts")
	dictRoutes.Use(middleware.JWTAuthMiddleware()) // 添加认证中间件
	{
		dictRoutes.GET("/", dictController.GetDictionaries)
		dictRoutes.POST("/", dictController.CreateDictionary)
		dictRoutes.PUT("/", dictController.UpdateDictionary)
		dictRoutes.DELETE("/:id", dictController.DeleteDictionary)
		dictRoutes.GET("/all", dictController.GetAllDictionaries)
	}
}
