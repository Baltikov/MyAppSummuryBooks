package routes

import (
	"log"
	"net/http"
	"strconv"
	"testapi/model"

	"github.com/gin-gonic/gin"
)

func getAllFAQ(context *gin.Context) {

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		log.Printf("Invalid limit parameter: %v", err)
		return
	}

	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		log.Printf("Invalid page parameter: %v", err)
		return
	}

	allFaq, err := model.GETAll(limit, page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Error getting FAQs: %v", err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": allFaq})
}
