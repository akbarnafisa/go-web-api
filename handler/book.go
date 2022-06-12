package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pongss",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	author := c.Param("author")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"author": author,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	titleArray := c.QueryArray("title")
	author := c.Query("author")

	c.JSON(http.StatusOK, gin.H{
		"title":      title,
		"titleArray": titleArray,
		"author":     author,
	})
}