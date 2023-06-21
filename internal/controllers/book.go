package controllers

import (
	"cloaks.cn/share/internal/models"
	"github.com/kataras/iris/v12"
)

type BookController struct {
	/* dependencies */
}

// GET: http://localhost:8080/books
func (c *BookController) Get() []models.Book {
	return []models.Book{
		{Title: "Mastering Concurrency in Go"},
		{Title: "Go Design Patterns"},
		{Title: "Black Hat Go"},
	}
}

// POST: http://localhost:8080/books
func (c *BookController) Post(b models.Book) int {
	println("Received Book: " + b.Title)

	return iris.StatusCreated
}
