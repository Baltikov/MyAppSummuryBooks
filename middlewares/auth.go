package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testapi/pkg/loger"
	"testapi/utils"
)

func Authentication(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")

	if authHeader == "" {
		loger.Logrus.Error("No Authorization Token")
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization Token"})
		return
	}

	// Проверяем наличие префикса "Bearer"
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		loger.Logrus.Error("Invalid Authorization Header")
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
		return
	}

	token := tokenParts[1]

	// Проверка токена
	userID, err := utils.CheckJwt(token)
	if err != nil {
		loger.Logrus.Error(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Token"})
		return
	}

	context.Set("UserID", userID)
	context.Next()
}
