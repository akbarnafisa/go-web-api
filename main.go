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
	bookHandler := handler.NewBookHandler((bookService))

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/books", bookHandler.GetBooks)
	v1.POST("/input", bookHandler.PostBookHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// main
	// handler
	// service
	// repository
	// db
	// mysql
}
