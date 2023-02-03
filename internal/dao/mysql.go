package dao

import (
	"trade-system/internal/model"
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
