package dao

import (
	"fmt"
	"strconv"
	"trade-system/internal/model"

	"gorm.io/gorm"
)

// SaveKLineInfo2Mysql 保存k线信息到mysql中
func (d *dao) SaveKLineInfo2Mysql(kLineInfo interface{}) error {
	switch v := kLineInfo.(type) {
	case model.KLineIn1Min:
		if err := d.MysqlDb.Create(&v).Error; err != nil {
			return err
		}
	case model.KLineIn5Min:
		if err := d.MysqlDb.Create(&v).Error; err != nil {
			return err
		}
	default:
	}
	return nil
}

func QueryKLineInfoFromMysql(d *dao, kLineType, amountStr string) (string, error) {
	amout, err := strconv.Atoi(amountStr)
	if err != nil {
		return "", fmt.Errorf("query amount invalid: %v", err)
	}
	switch kLineType {
	case "/kLine1Min":
		var kLineInfo []model.KLineIn1Min
		res := d.MysqlDb.Order("time desc").Limit(amout).Find(&kLineInfo)
		switch res.Error {
		case gorm.ErrInvalidDB:
			return "", fmt.Errorf("invalid db: %v", err)
		case gorm.ErrEmptySlice:
			return "", fmt.Errorf("data not found: %v", err)
		}
		if len(kLineInfo) == 0 {
			return "", fmt.Errorf("data not found")
		}
		return fmt.Sprintf("%+v", kLineInfo), nil
	case "/kLine5Min":
		var kLineInfo []model.KLineIn5Min
		res := d.MysqlDb.Order("time desc").Limit(amout).Find(&kLineInfo)
		switch res.Error {
		case gorm.ErrInvalidDB:
			return "", fmt.Errorf("invalid db: %v", err)
		case gorm.ErrEmptySlice:
			return "", fmt.Errorf("data not found: %v", err)
		}
		if len(kLineInfo) == 0 {
			return "", fmt.Errorf("data not found")
		}
		return fmt.Sprintf("%+v", kLineInfo), nil
	}
	return "", nil
}

func QueryKLineInfoFromMysqlWithKey(d *dao, kLineType, key, amountStr string) (string, error) {
	amout, err := strconv.Atoi(amountStr)
	if err != nil {
		return "", fmt.Errorf("query amount invalid: %v", err)
	}
	switch kLineType {
	case "kLine1Min":
		var kLineInfo []model.KLineIn1Min
		res := d.MysqlDb.Order(fmt.Sprintf("%s desc", key)).Limit(amout).Find(&kLineInfo)
		switch res.Error {
		case gorm.ErrInvalidDB:
			return "", fmt.Errorf("invalid db: %v", err)
		case gorm.ErrEmptySlice:
			return "", fmt.Errorf("data not found: %v", err)
		}
		if len(kLineInfo) == 0 {
			return "", fmt.Errorf("data not found")
		}
		return fmt.Sprintf("%+v", kLineInfo), nil
	case "kLine5Min":
		var kLineInfo []model.KLineIn5Min
		res := d.MysqlDb.Order(fmt.Sprintf("%s desc", key)).Limit(amout).Find(&kLineInfo)
		switch res.Error {
		case gorm.ErrInvalidDB:
			return "", fmt.Errorf("invalid db: %v", err)
		case gorm.ErrEmptySlice:
			return "", fmt.Errorf("data not found: %v", err)
		}
		if len(kLineInfo) == 0 {
			return "", fmt.Errorf("data not found")
		}
		return fmt.Sprintf("%+v", kLineInfo), nil

	}
	return "", nil
}
