package controllers

import (
	_ "cloaks.cn/share/docs"
	"cloaks.cn/share/pkg/configuration"
	"fmt"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/modrevision"
	"github.com/kataras/iris/v12/middleware/pprof"
	"github.com/kataras/iris/v12/mvc"
	"time"
)

func InitAPIs(api iris.Party) {
	api.PartyConfigure("/users", new(UserAPI))

	// mvc
	mvc.Configure(api.Party("/greet"), NewGreetController)
}

func InitCommonAPIs(api iris.Party, cfg configuration.Configuration) {
	// Add a simple health route.
	api.Any("/health", modrevision.New(modrevision.Options{
		ServerName:   cfg.ServerName,
		Env:          cfg.Env,
		Developer:    "cloaks",
		TimeLocation: time.FixedZone("Greece/Athens", 7200),
	}))

	// Configure the swagger UI page.
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL(fmt.Sprintf("http://%s:%d/swagger/doc.json", cfg.Host, cfg.Port)), // The url pointing to API definition.
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)
	// Register on http://localhost:8080/swagger
	api.Get("/swagger", swaggerUI)
	// And http://localhost:8080/swagger/index.html, *.js, *.css and e.t.c.
	api.Get("/swagger/{any:path}", swaggerUI)

	// http://localhost:8080/debug/pprof
	p := pprof.New()
	api.Any("/debug/pprof", p)
	api.Any("/debug/pprof/{action:path}", p)
}
