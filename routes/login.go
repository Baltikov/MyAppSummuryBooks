package routes

import (
	"net/http"
	"testapi/model"
	"testapi/pkg/loger"
	"testapi/utils"

	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		loger.Logrus.Error(err)
		loger.Logrus.Trace(err.Error())
	}
	err = model.Save(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
func login(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		loger.Logrus.Error(err)
		loger.Logrus.Trace(err.Error())
	}
	// обновить токен
	err = model.CheckUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		loger.Logrus.Error(err)
		loger.Logrus.Trace(err.Error())
		return
	}
	token, err := utils.GenerateJwt(user.Email, user.ID)
	if err != nil {
		loger.Logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// в заголовки прикрепить токен
	c.JSON(http.StatusOK, gin.H{"token": token})
}
