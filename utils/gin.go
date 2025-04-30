package utils

import (
	gin "github.com/gin-gonic/gin"
)

func SuccessResponseFormatter(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
	})
}

func ErrorResponseFormatter(c *gin.Context, statusCode int, message string, err error) {
	detailedMessage := ""
	if err != nil {
		detailedMessage = err.Error()
	}
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
		"error":   detailedMessage,
	})
}
