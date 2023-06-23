package controllers

import (
	"cloaks.cn/share/internal/models"
	"cloaks.cn/share/internal/services"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/mvc"
)

type GreetController struct {
	Logger  *golog.Logger
	Service services.GreetService
}

func NewGreetController(app *mvc.Application) {
	// Register Dependencies.
	app.Register(
		services.NewGreetService, // greeterWithLogging, greeter
	)

	// Register Controllers.
	app.Handle(new(GreetController))
}

func (c *GreetController) BeforeActivation(b mvc.BeforeActivation) {
	c.Logger = b.Dependencies().Logger

	// b.Dependencies().Add/Remove
	//b.Router().Use/UseGlobal/Done // 和你已知的任何标准 API  调用

	b.Handle("GET", "/sayHello", "SayHello")
}

// SayHello
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
func (c *GreetController) SayHello(req models.GreetReq) (models.GreetRes, error) {
	c.Logger.Infof("Request path: %s", "c.Ctx.Path()")
	message, err := c.Service.Say(req.Name)
	if err != nil {
		c.Logger.Warn("Shutdown with error: " + err.Error())
		return models.GreetRes{}, err
	}

	return models.GreetRes{Message: message}, nil
}
