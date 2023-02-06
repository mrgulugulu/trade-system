package server

import (
	"fmt"
	"net/http"
	"strings"
	"trade-system/config"
	"trade-system/internal/cache"
	"trade-system/internal/dao"
	"trade-system/internal/log"
	"trade-system/internal/model"

	"github.com/gin-gonic/gin"
)

// queryKLineIn1Min 查询一分钟k线的信息
func queryKLineIn1Min(c *gin.Context) {
	queryAmountStr := c.DefaultQuery("amount", "10")
	d, err := dao.NewDao()
	if err != nil {
		c.String(http.StatusInternalServerError, "database connect errors: %v", err)
		log.Sugar.Panicf("database connect errors: %v", err)
		return
	}
	// 先查询cache, key的格式为klinein1min+返回数量，如查询最新10条1分钟K线信息，则key为“klinein1min-10”
	cacheKey := fmt.Sprintf("%s-%s", config.CacheKeyKLineIn1Min, queryAmountStr)
	v, found := cache.C.Get(cacheKey)
	if found {
		res, err := v.([]model.KLineIn1Min)
		if err {
			c.String(http.StatusOK, fmt.Sprintf("%+v", res))
			log.Sugar.Infof("get data from cache successfully; data: ", res)
			return
		}
	}
	kLineType := c.Request.URL.Path
	res, err := dao.QueryKLineInfoFromMysql(d, kLineType, queryAmountStr)
	if err != nil {
		c.String(http.StatusInternalServerError, "query k line error: %v", err)
		log.Sugar.Errorf("query k line error: %v", err)
		return
	}
	cache.C.Set(cacheKey, res, config.CacheExpirationTime)
	c.String(http.StatusOK, fmt.Sprintf("%+v", res))
}

// queryKLineIn5Min 查询五分钟k线的信息
func queryKLineIn5Min(c *gin.Context) {
	queryAmountStr := c.DefaultQuery("amount", "10")
	d, err := dao.NewDao()
	if err != nil {
		c.String(http.StatusInternalServerError, "database connect errors: %v", err)
		log.Sugar.Panicf("database connect errors: %v", err)
		return
	}
	// 先查询cache, key的格式为klinein5min+返回数量，如查询最新10条5分钟K线信息，则key为“klinein5min10”
	cacheKey := fmt.Sprintf("%s-%s", config.CacheKeyKLineIn1Min, queryAmountStr)
	v, found := cache.C.Get(cacheKey)
	if found {
		res, err := v.([]model.KLineIn1Min)
		if err {
			c.String(http.StatusOK, fmt.Sprintf("%+v", res))
			log.Sugar.Infof("get data from cache successfully; data: ", res)
			return
		}
	}
	kLineType := c.Request.URL.Path
	res, err := dao.QueryKLineInfoFromMysql(d, kLineType, queryAmountStr)
	if err != nil {
		c.String(http.StatusInternalServerError, "query k line error: %v", err)
		log.Sugar.Errorf("query k line error: %v", err)
		return
	}
	cache.C.Set(cacheKey, res, config.CacheExpirationTime)
	c.String(http.StatusOK, fmt.Sprintf("%+v", res))
}

func queryKLineIn1MinWithKey(c *gin.Context) {
	key := c.Param("key")
	queryAmountStr := c.DefaultQuery("amount", "10")
	d, err := dao.NewDao()
	if err != nil {
		c.String(http.StatusInternalServerError, "database connect errors: %v", d)
		log.Sugar.Panicf("database connect errors: %v", err)
		return
	}

	// 先查询cache, key的格式为klinein1min+查询字段+返回数量，如查询10条最高的开盘价，则cache的key为“klinein1min-open-10”
	cacheKey := fmt.Sprintf("%s-%s-%s", config.CacheKeyKLineIn1Min, key, queryAmountStr)
	v, found := cache.C.Get(cacheKey)
	if found {
		res, err := v.([]model.KLineIn1Min)
		if err {
			c.String(http.StatusOK, fmt.Sprintf("%+v", res))
			log.Sugar.Infof("get data from cache successfully; data: ", res)
			return
		}
	}
	kLineType := strings.Split(c.Request.URL.Path, "/")[1]
	res, err := dao.QueryKLineInfoFromMysqlWithKey(d, kLineType, key, queryAmountStr)
	if err != nil {
		c.String(http.StatusInternalServerError, "query k line error: %v", err)
		log.Sugar.Errorf("query k line error: %v", err)
		return
	}

	cache.C.Set(cacheKey, res, config.CacheExpirationTime)
	c.String(http.StatusOK, fmt.Sprintf("%+v", res))
}

func queryKLineIn5MinWithKey(c *gin.Context) {
	key := c.Param("key")
	queryAmountStr := c.DefaultQuery("amount", "10")
	d, err := dao.NewDao()
	if err != nil {
		c.String(http.StatusInternalServerError, "database connect errors: %v", d)
		log.Sugar.Panicf("database connect errors: %v", err)
		return
	}

	// 先查询cache, key的格式为klinein1min+查询字段+返回数量，如查询10条最高的开盘价，则cache的key为“klinein1min-open-10”
	cacheKey := fmt.Sprintf("%s-%s-%s", config.CacheKeyKLineIn5Min, key, queryAmountStr)
	v, found := cache.C.Get(cacheKey)
	if found {
		res, err := v.([]model.KLineIn5Min)
		if err {
			c.String(http.StatusOK, fmt.Sprintf("%+v", res))
			log.Sugar.Infof("get data from cache successfully; data: ", res)
			return
		}
	}
	kLineType := strings.Split(c.Request.URL.Path, "/")[1]
	res, err := dao.QueryKLineInfoFromMysqlWithKey(d, kLineType, key, queryAmountStr)
	if err != nil {
		c.String(http.StatusInternalServerError, "query k line error: %v", err)
		log.Sugar.Errorf("query k line error: %v", err)
		return
	}
	cache.C.Set(cacheKey, res, config.CacheExpirationTime)
	c.String(http.StatusOK, fmt.Sprintf("%+v", res))
}
