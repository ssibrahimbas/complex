package wrapper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/tools/result"
)

func RouteWrapper() gin.HandlerFunc {
	return handleRouteErrors(gin.ErrorTypeAny)
}

func handleRouteErrors(errType gin.ErrorType) gin.HandlerFunc {
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
