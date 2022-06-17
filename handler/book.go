package handler

import (
	"go-web-api/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pongss",
	})
}

func (h *bookHandler) BookHandler(c *gin.Context) {
	id := c.Param("id")
	author := c.Param("author")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"author": author,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	titleArray := c.QueryArray("title")
	author := c.Query("author")

	c.JSON(http.StatusOK, gin.H{
		"title":      title,
		"titleArray": titleArray,
		"author":     author,
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {
	var bookInput book.BookRequest
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := h.bookService.Create(bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
