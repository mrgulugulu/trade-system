package server

import (
	"fmt"
	"net/http"
	"strconv"
	"trade-system/config"
	"trade-system/internal/dao"
	"trade-system/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

var C cache.Cache

func init() {
	C = *cache.New(config.CacheExpirationTime, config.CacheCleanUpInterval)
}

// queryKLineIn1Min 查询一分钟k线的信息
func queryKLineIn1Min(c *gin.Context) {
	queryAmountStr := c.DefaultQuery("amount", "10")
	d, err := dao.NewDao()
	if err != nil {
		c.String(http.StatusInternalServerError, "database connect errors: %v", d)
		return
	}
	// 先查询cache, key的格式为klinein1min+返回数量，如查询最新10条1分钟K线信息，则key为“klinein1min10”
	v, found := C.Get(config.CacheKeyKLineIn1Min + queryAmountStr)
	if found {
		res, err := v.([]model.KLineIn1Min)
		if err == true {
			c.String(http.StatusOK, fmt.Sprintf("%+v", res))
			return
		}
	}

	queryAmout, err := strconv.Atoi(queryAmountStr)
	if err != nil {
		c.String(http.StatusBadRequest, "query amount invalid: %v", err)
		return
	}

	var kLineInfo []model.KLineIn1Min
	res := d.MysqlDb.Order("time desc").Limit(queryAmout).Find(&kLineInfo)
	switch res.Error {
	case gorm.ErrInvalidDB:
		c.String(http.StatusInternalServerError, "invalid db: %v", err)
		return
	case gorm.ErrEmptySlice:
		c.String(http.StatusNotFound, "data not found: %v", err)
		return
	}
	if len(kLineInfo) == 0 {
		c.String(http.StatusNotFound, "data not found")
		return
	}
	// 设置缓存
	C.Set(config.CacheKeyKLineIn1Min+queryAmountStr, kLineInfo, config.CacheExpirationTime)
	c.String(http.StatusOK, fmt.Sprintf("%+v", kLineInfo))
}

// queryKLineIn5Min 查询一分钟k线的信息
func queryKLineIn5Min(c *gin.Context) {
	queryAmountStr := c.DefaultQuery("amount", "10")
	d, err := dao.NewDao()
	if err != nil {
		c.String(http.StatusInternalServerError, "database connect errors: %v", d)
		return
	}

	// 先查询cache, key的格式为klinein1min+返回数量，如查询最新10条1分钟K线信息，则key为“klinein1min10”
	v, found := C.Get(config.CacheKeyKLineIn5Min + queryAmountStr)
	if found {
		res, err := v.([]model.KLineIn1Min)
		if err == true {
			c.String(http.StatusOK, fmt.Sprintf("%+v", res))
			return
		}
	}
	queryAmout, err := strconv.Atoi(queryAmountStr)
	if err != nil {
		c.String(http.StatusBadRequest, "query amount invalid: %v", err)
		return
	}

	var kLineInfo []model.KLineIn5Min
	res := d.MysqlDb.Order("time desc").Limit(queryAmout).Find(&kLineInfo)
	switch res.Error {
	case gorm.ErrInvalidDB:
		c.String(http.StatusInternalServerError, "invalid db: %v", err)
	case gorm.ErrEmptySlice:
		c.String(http.StatusNotFound, "data not found: %v", err)
	}
	if len(kLineInfo) == 0 {
		c.String(http.StatusNotFound, "data not found")
		return
	}
	// 设置缓存
	C.Set(config.CacheKeyKLineIn5Min+queryAmountStr, kLineInfo, config.CacheExpirationTime)
	c.String(http.StatusOK, fmt.Sprintf("%+v", kLineInfo))
}
