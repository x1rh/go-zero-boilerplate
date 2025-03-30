package svc

import (
	"go-zero-boilerplate/internal/config"
	"go-zero-boilerplate/internal/model/db"
	"go-zero-boilerplate/pkg/zero-contrib/jwtx"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	JwtManager *jwtx.JWTManager

	MySQLConn sqlx.SqlConn
	// CacheDB   sqlc.CachedConn
	// BizCache *redis.Redis

	UserModel     db.UserModel
	TelegramModel db.TelegramModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MysqlDB.DataSource)
	return &ServiceContext{
		Config: c,

		JwtManager: jwtx.NewJwtManager(*c.JWT[0]),
		MySQLConn:  mysqlConn,

		// model
		UserModel:     db.NewUserModel(mysqlConn),
		TelegramModel: db.NewTelegramModel(mysqlConn),
	}
}
