package apis

import (
	"cloaks.cn/share/internal/controllers"
	"cloaks.cn/share/internal/models"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func BooksAPI(app *iris.Application) {
	booksAPI := app.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)
	}

	m := mvc.New(booksAPI)
	m.Handle(new(controllers.BookController))
}

func list(ctx iris.Context) {
	books := []models.Book{
		{Title: "Mastering Concurrency in Go"},
		{Title: "Go Design Patterns"},
		{Title: "Black Hat Go"},
	}

	ctx.JSON(books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

func create(ctx iris.Context) {
	var b models.Book
	err := ctx.ReadJSON(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
