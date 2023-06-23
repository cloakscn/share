package services

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"

	"cloaks.cn/share/pkg/database"
)

// GreetService example service.
type GreetService interface {
	Say(input string) (string, error)
}

// NewGreetService returns a service backed with a "db" based on "env".
func NewGreetService(db *database.DB, ctx iris.Context) GreetService {
	baseService := base{
		logger: ctx.Application().Logger(),
		db:     db,
	}
	return &greeter{
		base:   baseService,
		prefix: "Hello",
	}
}

type greeter struct {
	base
	prefix string
}

func (s *greeter) Say(input string) (string, error) {
	result := s.prefix + " " + input
	s.logger.Info(result)

	redis := s.db.Redis
	str, err := redis.Set(context.Background(), "say_hello", result, 0).Result()
	if err != nil {
		return "", err
	}

	return str, nil
}

type greeterWithLogging struct {
	*greeter
}

func (s *greeterWithLogging) Say(input string) (string, error) {
	result, err := s.greeter.Say(input)
	fmt.Printf("result: %s\nerror: %v\n", result, err)
	return result, err
}
