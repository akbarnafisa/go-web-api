package main

import (
	"go-web-api/book"
	"go-web-api/handler"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go-web-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connectionn error!")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	bookService := book.NewService(bookRepository)

	inputBook := book.BookRequest{
		Title: "Gundam",
		Price: "20000",
	}

	bookService.Create(inputBook)

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/ping", handler.RootHander)
	v1.GET("/books/:author/:id", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/input", handler.PostBookHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
