package svc

import (
	"go-zero-boilerplate/app/user-service/internal/config"
	"go-zero-boilerplate/pkg/zero-contrib/gormx"
	"gorm.io/gorm/logger"
)

func NewMysql(c config.Mysql) *gormx.DB {
	var logLevel logger.LogLevel
	switch c.LogLevel {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Info
	}

	db := gormx.MustNewMysql(&gormx.Config{
		DSN:          c.DataSource,
		MaxOpenConns: c.MaxOpenConns,
		MaxIdleConns: c.MaxIdleConns,
		MaxLifetime:  c.MaxLifetime,
		LogLevel:     logLevel,
	})
	return db
}