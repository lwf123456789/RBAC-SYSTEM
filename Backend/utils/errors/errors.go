package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIError 定义了 API 错误的结构
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Err        string `json:"error"`
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	return fmt.Sprintf("状态码: %d, 消息: %s, 错误: %s", e.StatusCode, e.Message, e.Err)
}

// NewAPIError 创建一个新的 APIError
func NewAPIError(statusCode int, message string, err error) *APIError {
	apiError := &APIError{
		StatusCode: statusCode,
		Message:    message,
	}
	if err != nil {
		apiError.Err = err.Error()
	}
	return apiError
}

// NewBadRequestError 创建一个 400 Bad Request 错误
func NewBadRequestError(message string, err error) *APIError {
	return NewAPIError(http.StatusBadRequest, message, err)
}

// NewUnauthorizedError 创建一个 401 Unauthorized 错误
func NewUnauthorizedError(message string, err error) *APIError {
	return NewAPIError(http.StatusUnauthorized, message, err)
}

// NewForbiddenError 创建一个 403 Forbidden 错误
func NewForbiddenError(message string, err error) *APIError {
	return NewAPIError(http.StatusForbidden, message, err)
}

// NewNotFoundError 创建一个 404 Not Found 错误
func NewNotFoundError(message string, err error) *APIError {
	return NewAPIError(http.StatusNotFound, message, err)
}

// NewInternalServerError 创建一个 500 Internal Server Error 错误
func NewInternalServerError(message string, err error) *APIError {
	return NewAPIError(http.StatusInternalServerError, message, err)
}

// HandleAPIError 处理 API 错误并发送适当的 JSON 响应
func HandleAPIError(c *gin.Context, err error) {
	apiError, ok := err.(*APIError)
	if !ok {
		apiError = NewInternalServerError("发生了未知错误", err)
	}

	c.JSON(apiError.StatusCode, gin.H{
		"error": apiError.Message,
	})
}

// IsNotFound 检查错误是否为 "Not Found" 错误
func IsNotFound(err error) bool {
	apiError, ok := err.(*APIError)
	return ok && apiError.StatusCode == http.StatusNotFound
}

// IsBadRequest 检查错误是否为 "Bad Request" 错误
func IsBadRequest(err error) bool {
	apiError, ok := err.(*APIError)
	return ok && apiError.StatusCode == http.StatusBadRequest
}

// IsUnauthorized 检查错误是否为 "Unauthorized" 错误
func IsUnauthorized(err error) bool {
	apiError, ok := err.(*APIError)
	return ok && apiError.StatusCode == http.StatusUnauthorized
}

// IsForbidden 检查错误是否为 "Forbidden" 错误
func IsForbidden(err error) bool {
	apiError, ok := err.(*APIError)
	return ok && apiError.StatusCode == http.StatusForbidden
}

// IsInternalServerError 检查错误是否为 "Internal Server Error" 错误
func IsInternalServerError(err error) bool {
	apiError, ok := err.(*APIError)
	return ok && apiError.StatusCode == http.StatusInternalServerError
}
