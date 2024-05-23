package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testapi/model"
)

func getAllFQA(context *gin.Context) {
	limitQuery := context.Query("limit")
	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{})
	}

	pageQuery := context.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{})
	}
	allFaq, err := model.GETAll(limit, page)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	context.JSON(http.StatusOK, gin.H{"data": allFaq})
}
