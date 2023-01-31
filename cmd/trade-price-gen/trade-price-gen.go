package main

import (
	"fmt"
	"math/rand"
	"trade-system/config"
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
	save2mysql(tradePairList)

}

// randFloats 生成n个指定范围的随机浮点数
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func save2mysql(tradePairList []model.TradePair) {
	config.LoadConfig()
	count := dao.D.InsertTradePair(tradePairList)
	fmt.Println(count)

}
