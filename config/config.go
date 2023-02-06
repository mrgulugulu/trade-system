// package config 项目的配置参数设置
package config

import (
	"time"
	"trade-system/internal/log"

	"github.com/spf13/viper"
)

var (
	// 交易对chan
	TradePairChannelName = "trade_pair"
	// 用于订阅redis的channel名
	KLineIn1MinChannelName = "1min_k_line"
	KLineIn5MinChannelName = "5min_k_line"
	// cache过期时间
	CacheExpirationTime = 30 * time.Second
	// cache清理间隔
	CacheCleanUpInterval = 60 * time.Second
	// 1分钟k线cache的前半部分key
	CacheKeyKLineIn1Min = "klinein1min"
	// 5分钟k线cache的前半部分key
	CacheKeyKLineIn5Min = "klinein5min"
)

type MysqlConfig struct {
	MysqlIP   string `mapstructure:"ip"`
	MysqlPort string `mapstructure:"port"`
	MysqlUser string `mapstructure:"user"`
	MysqlPwd  string `mapstructure:"pwd"`
	DataBase  string `mapstructure:"database"`
}

type RedisConfig struct {
	RedisIP   string `mapstructure:"ip"`
	RedisPort string `mapstructure:"port"`
	RedisPwd  string `mapstructure:"pwd"`
	DataBase  int    `mapstructure:"db"`
}
type Config struct {
	Mysqlcfg  *MysqlConfig  `mapstructure:"mysql"`
	Rediscfg  *RedisConfig  `mapstructure:"redis"`
	ServerCfg *ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	Addr string `mapstructure:"addr"`
	Port string `mapstructure:"port"`
}

var ServiceConf *Config

// LoadConfig 载入设置
func LoadConfig(confFile ...string) {
	c := viper.New()
	conf := Config{}
	c.AddConfigPath("../../config")
	c.AddConfigPath("../config")
	c.SetConfigName("config")
	c.SetConfigType("yaml")
	err := c.ReadInConfig()
	if err != nil {
		log.Sugar.Infof("read config error: %v", err)
	}
	err = c.Unmarshal(&conf)
	ServiceConf = &conf
}
