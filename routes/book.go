package routes

import (
	"fmt"
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

	// Создать может только авторизованный пользователь

	// Мы понимаем, что айди пользователя мы можешь получить в токене
	// возвращаем его и записываем в таблицу.
	// таким образом, при создании мы указываем какой пользователь создал эту запись

	var modelBook model.Book
	err := context.ShouldBind(&modelBook)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	userID := context.GetInt64("UserID")
	modelBook.UserID = userID

	err = model.CreateBook(modelBook)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"data": modelBook})

}
func updateBook(context *gin.Context) {
	idParam := context.Param("id")
	_, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var modelBook model.Book
	err = context.ShouldBind(&modelBook)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	userID := context.GetInt64("UserID")
	// обновить стать может пользователь если его айди совпало с айди ключа
	if modelBook.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Token"})
	}

	updateBooks, err := model.UpdateBook(modelBook)
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
	userID := context.GetInt64("UserID")

	modelBook, err := model.GetBook(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if modelBook.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Token"})
	}
	err = model.DeleteBook(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Удалена книга с айди %s", modelBook.ID)})

}
