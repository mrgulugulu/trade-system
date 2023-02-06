package server

import (
	"fmt"
	"net/http"
	"strconv"
	"trade-system/config"
	"trade-system/internal/cache"
	"trade-system/internal/dao"
	"trade-system/internal/log"
	"trade-system/internal/model"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
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
	// 先查询cache, key的格式为klinein1min+返回数量，如查询最新10条1分钟K线信息，则key为“klinein1min10”
	v, found := cache.C.Get(config.CacheKeyKLineIn1Min + queryAmountStr)
	if found {
		res, err := v.([]model.KLineIn1Min)
		if err == true {
			c.String(http.StatusOK, fmt.Sprintf("%+v", res))
			log.Sugar.Infof("get data from cache successfully; data: ", res)
			return
		}
	}

	queryAmout, err := strconv.Atoi(queryAmountStr)
	if err != nil {
		c.String(http.StatusBadRequest, "query amount invalid: %v", err)
		log.Sugar.Errorf("query amount invalid: %v", err)
		return
	}

	var kLineInfo []model.KLineIn1Min
	res := d.MysqlDb.Order("time desc").Limit(queryAmout).Find(&kLineInfo)
	switch res.Error {
	case gorm.ErrInvalidDB:
		c.String(http.StatusInternalServerError, "invalid db: %v", err)
		log.Sugar.Errorf("invalid db: %v", err)
		return
	case gorm.ErrEmptySlice:
		c.String(http.StatusNotFound, "data not found: %v", err)
		log.Sugar.Errorf("data not found: %v", err)
		return
	}
	if len(kLineInfo) == 0 {
		c.String(http.StatusNotFound, "data not found")
		log.Sugar.Errorf("data not found")
		return
	}
	// 设置缓存
	cache.C.Set(config.CacheKeyKLineIn1Min+queryAmountStr, kLineInfo, config.CacheExpirationTime)
	c.String(http.StatusOK, fmt.Sprintf("%+v", kLineInfo))
}

// queryKLineIn5Min 查询一分钟k线的信息
func queryKLineIn5Min(c *gin.Context) {
	queryAmountStr := c.DefaultQuery("amount", "10")
	d, err := dao.NewDao()
	if err != nil {
		c.String(http.StatusInternalServerError, "database connect errors: %v", d)
		log.Sugar.Panicf("database connect errors: %v", err)
		return
	}

	// 先查询cache, key的格式为klinein1min+返回数量，如查询最新10条1分钟K线信息，则key为“klinein1min10”
	v, found := cache.C.Get(config.CacheKeyKLineIn5Min + queryAmountStr)
	if found {
		res, err := v.([]model.KLineIn1Min)
		if err == true {
			c.String(http.StatusOK, fmt.Sprintf("%+v", res))
			log.Sugar.Infof("get data from cache successfully; data: ", res)
			return
		}
	}
	queryAmout, err := strconv.Atoi(queryAmountStr)
	if err != nil {
		c.String(http.StatusBadRequest, "query amount invalid: %v", err)
		log.Sugar.Errorf("query amount invalid: %v", err)
		return
	}

	var kLineInfo []model.KLineIn5Min
	res := d.MysqlDb.Order("time desc").Limit(queryAmout).Find(&kLineInfo)
	switch res.Error {
	case gorm.ErrInvalidDB:
		c.String(http.StatusInternalServerError, "invalid db: %v", err)
		log.Sugar.Errorf("invalid db: %v", err)
	case gorm.ErrEmptySlice:
		c.String(http.StatusNotFound, "data not found: %v", err)
		log.Sugar.Errorf("data not found: %v", err)
	}
	if len(kLineInfo) == 0 {
		c.String(http.StatusNotFound, "data not found")
		log.Sugar.Errorf("data not found")
		return
	}
	// 设置缓存
	cache.C.Set(config.CacheKeyKLineIn5Min+queryAmountStr, kLineInfo, config.CacheExpirationTime)
	c.String(http.StatusOK, fmt.Sprintf("%+v", kLineInfo))
}
