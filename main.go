package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", rootHander)
	r.GET("/books/:author/:id", bookHandler)
	r.GET("/query", queryHandler)
	r.POST("/input", postBookHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func rootHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pongss",
	})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	author := c.Param("author")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"author": author,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	titleArray := c.QueryArray("title")
	author := c.Query("author")

	c.JSON(http.StatusOK, gin.H{
		"title":      title,
		"titleArray": titleArray,
		"author":     author,
	})
}

type BookInput struct {
	Title    string
	Price    int
	SubTitle string `json:"sub_title"`
}

func postBookHandler(c *gin.Context) {
	var bookInput BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "Books added",
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
