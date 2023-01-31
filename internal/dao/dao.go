package dao

import (
	"fmt"
	"trade-system/config"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var D *Dao

func init() {
	config.LoadConfig()
	D, _ = newDao(config.ServiceConf)
}

type Dao struct {
	Config  *config.Config
	MysqlDb *gorm.DB
	RedisDb *redis.Client
}

// newDao 初始化Dao
func newDao(c *config.Config) (*Dao, error) {
	// 连接mysql
	mysqlDb, err := newMysql(c.Mysqlcfg)
	if err != nil {
		return nil, err
	}
	redisDb, err := newRedis(c.Rediscfg)
	if err != nil {
		return nil, err
	}
	dao := &Dao{
		MysqlDb: mysqlDb,
		RedisDb: redisDb,
		Config:  c,
	}
	return dao, nil

}

// newMysql 初始化Mysql
func newMysql(c *config.MysqlConfig) (*gorm.DB, error) {
	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?&charset=utf8mb4&parseTime=True&loc=Local",
		c.MysqlUser, c.MysqlPwd, c.MysqlIP, c.MysqlPort, c.DataBase)
	db, err := gorm.Open(mysql.Open(connArgs), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// newRedis初始化Redis
func newRedis(c *config.RedisConfig) (*redis.Client, error) {
	redisOptions := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.RedisIP, c.RedisPort),
		Password: c.RedisPwd,
		DB:       c.DataBase,
	}
	redisDb := redis.NewClient(redisOptions)
	// 测试链接
	if _, err := redisDb.Ping().Result(); err != nil {
		return nil, err
	}

	return redisDb, nil
}
