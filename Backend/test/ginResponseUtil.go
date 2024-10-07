package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一响应结构体
type Response struct {
	Code    int         `json:"code"`    // HTTP 状态码
	Message string      `json:"message"` // 消息
	Data    interface{} `json:"data"`    // 数据
}

// 成功响应封装
func SuccessRes(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// 错误响应封装
func ErrorRes(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func main() {
	r := gin.Default()

	// 示例 API 路由
	r.GET("/products", getProducts)
	r.GET("/products/:id", getProductByID)

	r.Run(":8080")
}

// 模拟的商品数据
var products = []map[string]interface{}{
	{"id": "1", "name": "Product A", "price": 100},
	{"id": "2", "name": "Product B", "price": 200},
}

// 获取所有商品
func getProducts(c *gin.Context) {
	SuccessRes(c, products)
}

// 根据 ID 获取商品
func getProductByID(c *gin.Context) {
	id := c.Param("id")
	for _, product := range products {
		if product["id"] == id {
			SuccessRes(c, product)
			return
		}
	}
	// 如果商品未找到，返回错误响应
	ErrorRes(c, http.StatusNotFound, "Product not found")
}
