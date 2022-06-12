package main

import (
	"go-web-api/book"
	"go-web-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/ping", handler.RootHander)
	v1.GET("/books/:author/:id", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/input", book.PostBookHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
