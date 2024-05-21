package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"testapi/bd"
)

func main() {
	bd.InitDB()
	defer bd.DB.Close()
	server := gin.Default()
	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
