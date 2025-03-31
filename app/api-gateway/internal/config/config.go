package config

import (
	"go-zero-boilerplate/pkg/zero-contrib/jwtx"

	"github.com/zeromicro/go-zero/gateway"
)

type Config struct {
	Gateway gateway.GatewayConf
	JWT []*jwtx.Config `json:"JWT"`
	MySQL   Mysql     `json:"MySQLConf"`
	Redis   RedisConf `json:"RedisConf"`
}

type Mysql struct {
	DataSource   string `json:"DataSource"`
	MaxOpenConns int    `json:"MaxOpenConns"`
	MaxIdleConns int    `json:"MaxIdleConns"`
	MaxLifetime  int    `json:"MaxLifetime"`
	LogLevel     string `json:"LogLevel"`
}

type RedisConf struct {
	Addr         string   `mapstructure:"addr"         json:"addr"         yaml:"addr"`         // 服务器地址:端口
	Password     string   `mapstructure:"password"     json:"password"     yaml:"password"`     // 密码
	DB           int      `mapstructure:"db"           json:"db"           yaml:"db"`           // 单实例模式下redis的哪个数据库
	UseCluster   bool     `mapstructure:"useCluster"   json:"useCluster"   yaml:"useCluster"`   // 是否使用集群模式
	ClusterAddrs []string `mapstructure:"clusterAddrs" json:"clusterAddrs" yaml:"clusterAddrs"` // 集群模式下的节点地址列表
}