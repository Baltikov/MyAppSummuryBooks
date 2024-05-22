package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/books", getBooks)
	server.POST("/books", createBook)
	server.PUT("/books/:id", updateBook)
	server.DELETE("/books/:id", deleteBook)
}
