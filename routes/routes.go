package routes

import (
	"github.com/gin-gonic/gin"
	"testapi/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	auth := server.Group("/books")
	auth.Use(middlewares.Authentication)
	auth.GET("/books", getBooks)
	auth.POST("/books", createBook)
	auth.PUT("/books/:id", updateBook)
	auth.DELETE("/books/:id", deleteBook)
	// server.POST("/books", middlewares.Authentication, createBook)

	server.GET("FAQ", getAllFAQ)
	server.POST("signUp", signUp)
	server.POST("login", login)
}
