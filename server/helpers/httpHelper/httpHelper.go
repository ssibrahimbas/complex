package httpHelper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/core/result"
)

func ErrorWrapper() gin.HandlerFunc {
	return handleErrors(gin.ErrorTypeAny)
}

func handleErrors(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)
		if len(detectedErrors) > 0 {
			c.JSON(http.StatusInternalServerError, result.ErrorData("An error occurred.", detectedErrors))
			c.Abort()
			return
		}
	}
}