package middleware

import (
	"GOCOM/models"
	"GOCOM/utils/errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const SecretKey = "123456" // 注意: 在实际应用中,应该使用环境变量存储密钥

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errors.HandleAPIError(c, errors.NewUnauthorizedError("需要授权头信息", nil))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			errors.HandleAPIError(c, errors.NewUnauthorizedError("授权头格式必须为Bearer {token}", nil))
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := &models.CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SecretKey), nil
		})

		if err != nil {
			errors.HandleAPIError(c, errors.NewUnauthorizedError("无效或者过期的token", err))
			c.Abort()
			return
		}

		if !token.Valid {
			errors.HandleAPIError(c, errors.NewUnauthorizedError("无效的token", nil))
			c.Abort()
			return
		}
		// 将用户信息保存到上下文中
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
