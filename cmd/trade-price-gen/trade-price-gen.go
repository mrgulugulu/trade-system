package main

import (
	"fmt"
	"log"
	"math/rand"
	"trade-system/internal/dao"
	"trade-system/internal/model"
)

var (
	minPrice = 20000.0
	maxPrice = 25000.0

	minAmt = 0.0001
	maxAmt = 1.0

	minDiff = 0.0
	maxDiff = 2.0

	genAmount = 10
)

func main() {
	d, err := dao.NewDao()
	if err != nil {
		fmt.Printf("new dao error: %v", err)
		return
	}
	priceList := randFloats(minPrice, maxPrice, genAmount)
	amtList := randFloats(minAmt, maxAmt, genAmount)
	// total >= amt
	diffList := randFloats(minDiff, maxDiff, genAmount)
	totalList := make([]float64, genAmount)
	tradePairList := make([]model.TradePair, genAmount)
	for i := 0; i < genAmount; i++ {
		totalList[i] = amtList[i] + diffList[i]
		tradePairList[i] = model.TradePair{Price: priceList[i], Amt: amtList[i], Total: totalList[i]}
	}

	count := 0
	for i, p := range tradePairList {
		if err := d.MysqlDb.Create(&p).Error; err != nil {
			log.Printf("db.Create index: %d, err : %v", i, err)
			continue
		}
		count++
	}

	fmt.Printf("成功插入%d条数据", count)

}

// randFloats 生成n个指定范围的随机浮点数
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}
