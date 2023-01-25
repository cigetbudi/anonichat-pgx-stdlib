package middlewares

import (
	"anonichat-pgx-stdlib/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.TokenValid(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
