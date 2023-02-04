package server

import (
	"fmt"
	"net/http"
	"trade-system/internal/dao"
	"trade-system/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// queryKLineIn1Min 查询一分钟k线的信息
func queryKLineIn1Min(c *gin.Context) {
	d, err := dao.NewDao()
	if err != nil {
		c.String(http.StatusInternalServerError, "databases connect errors: %v", d)
		return
	}
	var kLineInfo []model.KLineIn1Min
	res := d.MysqlDb.Order("time desc").Limit(10).Find(&kLineInfo)
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
	c.String(http.StatusOK, fmt.Sprintf("%+v", kLineInfo))
}

// queryKLineIn5Min 查询一分钟k线的信息
func queryKLineIn5Min(c *gin.Context) {
	d, err := dao.NewDao()
	if err != nil {
		c.String(http.StatusInternalServerError, "databases connect errors: %v", d)
		return
	}
	var kLineInfo []model.KLineIn5Min
	res := d.MysqlDb.Order("time desc").Limit(10).Find(&kLineInfo)
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
	c.String(http.StatusOK, fmt.Sprintf("%+v", kLineInfo))
}
