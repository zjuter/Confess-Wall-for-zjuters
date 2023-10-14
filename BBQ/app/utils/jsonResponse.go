package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, httpStatusCode int, code int, message string, data interface{}) {
	c.JSON(httpStatusCode, gin.H{
		"code": code,
		"message":  message,
	})
}

func JsonErrorResponse(c *gin.Context, code int, message string) {
	JsonResponse(c, http.StatusOK, code, message, nil)
}

func JsonInternalServerErrorResponse(c *gin.Context) {
	JsonResponse(c,http.StatusOK, 200500, "Internal server error", nil)
}
