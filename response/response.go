package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, APIResponse{
		Success: false,
		Error:   message,
	})
}

func BadRequestError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message)
}

func UnauthorizedError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message)
}

func InternalServerError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message)
}

func CreatedResponse(c *gin.Context, message string, data interface{}) {
	SuccessResponse(c, http.StatusCreated, message, data)
}

func OKResponse(c *gin.Context, message string, data interface{}) {
	SuccessResponse(c, http.StatusOK, message, data)
}
