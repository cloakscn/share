package main

import (
	"cloaks.cn/share/internal/apis"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	apis.BooksAPI(app)

	app.Listen(":8080")
}


