package services

import (
	"cloaks.cn/share/pkg/database"
	"github.com/kataras/golog"
)

type base struct {
	logger *golog.Logger
	db     *database.DB
}
