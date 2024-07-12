package middleware

import (
	auth "event_scheduler/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IsLogin Authorize the user
func IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unable to login",
			})
			ctx.Abort()
			return
		}
		claims, err := auth.ValidateToken(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("id", claims.UserID)
		ctx.Next()
	}
}
