package svc

import (
	"go-zero-boilerplate/app/api-gateway/internal/config"
	// "go-zero-boilerplate/pkg/zero-contrib/gormx"
	// "go-zero-boilerplate/pkg/zero-contrib/jwtx"

	"github.com/x1rh/pkg/zero-contrib/jwtx"
	"github.com/x1rh/pkg/zero-contrib/gormx"

	"github.com/redis/go-redis/v9"
)

type ServiceContext struct {	
	Config     config.Config
	DB          *gormx.DB
	Redis       redis.UniversalClient
	JwtManager *jwtx.Manager
}

func NewServiceContext(c config.Config) *ServiceContext {
	// rds := MustNewRedis(c.Redis)
	return &ServiceContext{
		Config: c,
		JwtManager: jwtx.NewManager(c.JWT),
		DB: NewMysql(c.MySQL),
		// Redis: rds, 
	}
}
