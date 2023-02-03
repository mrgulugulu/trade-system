// package kline 对k线进行相关的处理
package kline

import (
	"math"
	"trade-system/internal/model"
)

// KLineIn1MinGen 生成1分钟k线
func KLineIn1MinGen(tradePairList []model.TradePairWithTime) model.KLineIn1Min {
	// 时间切片的划分交给其他函数来处理
	// 需要记录以下字段的信息：开盘，收盘，最高价，最低价与交易量
	var op, cl, vol float64
	var hPrice = -math.MaxFloat64 - 1.0
	var lPrice = math.MaxFloat64
	op, cl = tradePairList[0].Price, tradePairList[len(tradePairList)-1].Price
	for _, p := range tradePairList {
		if p.Price > hPrice {
			hPrice = p.Price
		} else if p.Price < lPrice {
			lPrice = p.Price
		}
		vol += p.Total
	}
	return model.KLineIn1Min{
		// 12:01分的一分钟k线指12:00:00-12:00:59的k线信息
		Time:         tradePairList[len(tradePairList)-1].Time,
		Open:         op,
		Close:        cl,
		HighestPrice: hPrice,
		LowestPrice:  lPrice,
		Volume:       vol,
	}
}

// KLineIn5MinGen 生成5分钟k线
func KLineIn5MinGen(tradePairList []model.TradePairWithTime) model.KLineIn5Min {
	// 时间切片的划分交给其他函数来处理
	// 需要记录以下字段的信息：开盘，收盘，最高价，最低价与交易量
	var op, cl, vol float64
	var hPrice = -math.MaxFloat64 - 1.0
	var lPrice = math.MaxFloat64
	op, cl = tradePairList[0].Price, tradePairList[len(tradePairList)-1].Price
	for _, p := range tradePairList {
		if p.Price > hPrice {
			hPrice = p.Price
		} else if p.Price < lPrice {
			lPrice = p.Price
		}
		vol += p.Total
	}
	return model.KLineIn5Min{
		// 12:05分的一分钟k线指12:00:00-12:04:49的k线信息
		Time:         tradePairList[len(tradePairList)-1].Time,
		Open:         op,
		Close:        cl,
		HighestPrice: hPrice,
		LowestPrice:  lPrice,
		Volume:       vol,
	}
}
