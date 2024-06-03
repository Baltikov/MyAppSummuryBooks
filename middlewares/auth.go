package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testapi/pkg/loger"
	"testapi/utils"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		loger.Logrus.Error("No Authorization Token")
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization Token"})
	}
	// Создать может только авторизованный пользователь
	userID, err := utils.CheckJwt(token)
	if err != nil {
		loger.Logrus.Error(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Token"})
	}
	context.Set("UserID", userID)
	context.Next()

}
