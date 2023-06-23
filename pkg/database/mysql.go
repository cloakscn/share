package database

import (
	"cloaks.cn/share/pkg/configuration"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mysqlOpen(cfg configuration.MysqlConfiguration) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		cfg.UserName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DataBase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
