package middlewares

import (
	"anonichat-pgx-stdlib/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.TokenValid(ctx)
		if err != nil {
			utils.RetUnauth(ctx, err)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
