package controllers

import (
	"time"

	"cloaks.cn/share/internal/models"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/mvc"
)

type TestController struct {
	Logger  *golog.Logger
}

func NewTestController(app *mvc.Application) {
	// Register Dependencies.
	app.Register(
	)

	// Register Controllers.
	app.Handle(new(TestController))
}

func (c *TestController) BeforeActivation(b mvc.BeforeActivation) {
	c.Logger = b.Dependencies().Logger

	// b.Dependencies().Add/Remove
	//b.Router().Use/UseGlobal/Done // 和你已知的任何标准 API  调用

	b.Handle("GET", "/post", "Post")
}

// Post
// @Description say hello
// @Tags v1
// @Accept json
// @Produce json
// @Param some_id path string true "Some ID"
// @Param offset query int true "Offset"
// @Param limit query int true "Offset"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} models.APIError "We need ID!!"
// @Failure 404 {object} models.APIError "Can not find ID"
// @Router /greet/sayHello [get]
func (c *TestController) Post(req models.GreetReq) (models.GreetRes, error) {
	c.Logger.Infof("Request path: %s", "c.Ctx.Path()")
	
	return models.GreetRes{Message: time.Now().GoString()}, nil
}
