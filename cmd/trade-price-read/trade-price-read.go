package main

import (
	"encoding/json"
	"fmt"
	"time"
	"trade-system/config"
	"trade-system/internal/dao"
	"trade-system/internal/model"
)

func main() {
	d, err := dao.NewDao()
	if err != nil {
		fmt.Printf("new dao error: %v", err)
		return
	}

	rows, err := d.MysqlDb.Model(&model.TradePair{}).Rows()
	if err != nil {
		fmt.Printf("mysql query error: %v", err)
		return
	}

	defer rows.Close()
	var count int
	// 通过迭代读出来
	for rows.Next() {

		var tradePair model.TradePair
		d.MysqlDb.ScanRows(rows, &tradePair)
		tradePairByte, err := json.Marshal(tradePair)
		if err != nil {
			fmt.Printf("data %v cannot be marshaled: %v", tradePair, err)
		}
		tradePairStr := string(tradePairByte)
		_, err = d.RedisDb.Publish(config.TradePairChannel, tradePairStr).Result()
		if err != nil {
			fmt.Printf("trade-pair: %v publish error: %v", tradePairStr, err)
		}
		count++
		if count == 9 {
			time.Sleep(time.Second)
			count = 0
		}
	}
}
