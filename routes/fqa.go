package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testapi/model"
)

func getAllFQA(context *gin.Context) {
	allFaq, err := model.GETAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	context.JSON(http.StatusOK, gin.H{"data": allFaq})
}
