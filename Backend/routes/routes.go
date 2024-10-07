package routes

import (
	"GOCOM/routes/admin"
	"GOCOM/routes/department"
	"GOCOM/routes/dictionary"
	"GOCOM/routes/menu"
	"GOCOM/routes/role"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 配置 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"}, // 前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true, // 允许跨域发送凭证，如 Cookies
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 60 * 60, // 预检请求结果缓存时间（秒）
	}))

	// 注册各模块路由
	department.RegisterDepartmentRoutes(r)
	admin.RegisterAdminRoutes(r)
	dictionary.RegisterDictRoutes(r)
	role.RegisterRoleRoutes(r)
	menu.RegisterMenuRoutes(r)
	return r
}
