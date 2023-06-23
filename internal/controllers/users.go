package controllers

import (
	"cloaks.cn/share/internal/models"
	"cloaks.cn/share/pkg/database"
	"cloaks.cn/share/pkg/http/handlers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/auth"
)

type UserAPI struct {
	DB   *database.DB // exported field so api/router.go#api.RegisterDependency can bind it.
	Auth *auth.Auth[models.User]
}

func (api *UserAPI) Configure(r iris.Party) {
	r.Post("/signup", api.signUp)
	r.Post("/signin", api.signIn)
	// Add middlewares such as user verification by bearer token here.

	// Authenticated routes...
	r.Get("/", api.getInfo)

	// GET http://localhost:8080/api/signin
	r.Get("/signin", api.Auth.SigninHandler)
	// GET http://localhost:8080/api/refresh
	r.Get("/refresh", api.Auth.RefreshHandler)
	// GET http://localhost:8080/api/signout
	r.Get("/signout", api.Auth.SignoutHandler)
	// GET http://localhost:8080/api/member
	r.Get("/member", api.Auth.VerifyHandler(handlers.AllowRole(models.Member)), func(s *auth.Auth[models.User]) iris.Handler {
		return func(ctx iris.Context) {
			user := s.GetUser(ctx)
			ctx.Writef("Hello member: %s\n", user.Email)
		}
	}(api.Auth))
	// GET http://localhost:8080/api/owner
	r.Get("/owner", api.Auth.VerifyHandler(handlers.AllowRole(models.Owner)), func(s *auth.Auth[models.User]) iris.Handler {
		return func(ctx iris.Context) {
			user := s.GetUser(ctx)
			ctx.Writef("Hello owner: %s\n", user.Email)
		}
	}(api.Auth))
}

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Router /testapi/get-string-by-int/{some_id} [get]

func (api *UserAPI) getInfo(ctx iris.Context) {
	ctx.WriteString("...")
}

func (api *UserAPI) signUp(ctx iris.Context) {}
func (api *UserAPI) signIn(ctx iris.Context) {}
