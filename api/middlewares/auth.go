package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodolfoksveiga/k8s-go/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := utils.ValidateToken(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
