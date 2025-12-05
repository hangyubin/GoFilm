package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ErrorResponse 统一错误响应格式
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data,omitempty"`
}

// ErrorHandler 全局错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录错误日志
				log.Printf("Panic recovered: %v", err)
				
				// 返回统一的错误响应
				c.JSON(http.StatusInternalServerError, ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "服务器内部错误",
				})
				
				// 终止后续中间件和处理函数的执行
				c.Abort()
			}
		}()
		
		c.Next()
		
		// 检查是否有错误
		if len(c.Errors) > 0 {
			// 获取最后一个错误
			err := c.Errors.Last()
			log.Printf("Request error: %v", err.Error())
			
			// 根据错误类型返回不同的状态码
			statusCode := http.StatusInternalServerError
			message := "服务器内部错误"
			
			// 可以根据错误类型进行更细粒度的处理
			
			// 返回统一的错误响应
			c.JSON(statusCode, ErrorResponse{
				Code:    statusCode,
				Message: message,
			})
			
			// 终止后续中间件的执行
			c.Abort()
		}
	}
}
