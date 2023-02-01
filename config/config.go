package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	TradePairChannel   = "trade_pair"
	KLineIn1MinChannel = "1min_k_line"
	KLineIn5MinChannel = "5min_k_line"
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

func LoadConfig(confFile ...string) error {
	c := viper.New()
	conf := Config{}
	c.AddConfigPath("../../config")
	c.AddConfigPath("./config")
	c.SetConfigName("config")
	c.SetConfigType("yaml")
	err := c.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read config error: %v", err)
	}
	err = c.Unmarshal(&conf)
	ServiceConf = &conf
	return err
}
