package book

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

func PostBookHandler(c *gin.Context) {
	var bookInput BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Books added",
		"title":  bookInput.Title,
		"price":  bookInput.Price,
	})
}
