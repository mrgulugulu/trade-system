package dao

import (
	"log"
	"trade-system/internal/model"
)

func (d *Dao) InsertTradePair(tradePrice []model.TradePair) int {
	count := 0
	for i, p := range tradePrice {
		if err := d.MysqlDb.Create(&p).Error; err != nil {
			log.Printf("db.Create index: %d, err : %v", i, err)
			continue
		}
		count++
	}
	return count
}
