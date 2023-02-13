package middlewares

import (
	"api/src/internal/pkg/jwt"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticatedGuard(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	err := jwt.ValidateToken(token)

	fmt.Println(err)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
