package middleware

import (
	auth "event_scheduler/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized access",
			})
			ctx.Abort()
			return
		}
		claims, err := auth.ValidateToken(tokenString)
		fmt.Println(err)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			ctx.Abort()
			return
		}
		ctx.Set("id", claims.Id)
		ctx.Next()
	}
}
