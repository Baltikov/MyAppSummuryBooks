package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testapi/model"
)

func getBooks(context *gin.Context) {
	books, err := model.GetBooks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": books})
}
