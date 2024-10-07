package main

import (
	"GOCOM/config"
	_ "GOCOM/docs" // 加载swagger 文档
	"GOCOM/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GOCOM API
// @version 1.0
// @description 这是GOCOM项目的API文档，提供了详细的接口说明和使用方法。
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.yourcompany.com/support
// @contact.email support@yourcompany.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8082
// @BasePath
// @schemes http https

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description 请输入 'Bearer {token}' 来进行身份验证

func main() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化数据库
	config.InitDB()
	defer config.CloseDB()

	// 设置路由
	r := routes.SetupRouter()

	// 提供静态文件服务，映射 /static/images 目录到 /images 路径
	r.StaticFS("/images", http.Dir("./static/images"))
	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 启动服务器
	r.Run(":8082")

	// 阻止主函数退出
	select {}
}
