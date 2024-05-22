package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testapi/model"
)

func getBooks(context *gin.Context) {
	books, err := model.GetBooks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": books})
}
func createBook(context *gin.Context) {

	var modelBook model.Book
	err := context.ShouldBind(&modelBook)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	books, err := model.CreateBook(modelBook)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"data": books})

}
func updateBook(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	var modelBook model.Book
	err = context.ShouldBind(&modelBook)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	updateBooks, err := model.UpdateBook(modelBook, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": updateBooks})
}
func deleteBook(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	books, err := model.DeleteBook(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": books})
}
