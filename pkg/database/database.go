package database

import (
	"cloaks.cn/share/pkg/configuration"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DB struct {
	Mysql *gorm.DB
	Redis *redis.Client
}

func NewDB(cfg configuration.DBConfiguration) *DB {
	db := &DB{}
	// init mysql.
	if cfg.Mysql != nil {
		db.Mysql = mysqlOpen(*cfg.Mysql)
	}
	// init sqlite.
	if cfg.Sqlite != nil {

	}
	// init redis.
	if cfg.Redis != nil {
		db.Redis = redisOpen(*cfg.Redis)
	}

	return db
}
