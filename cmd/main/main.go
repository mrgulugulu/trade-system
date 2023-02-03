package main

import (
	"encoding/json"
	"fmt"
	"trade-system/config"
	"trade-system/internal/dao"
	"trade-system/internal/kline"
	"trade-system/internal/model"
)

func main() {
	d, err := dao.NewDao()
	if err != nil {
		fmt.Printf("new dao error: %v", err)
		return
	}
	// 需求
	// 1. 从redis中订阅交易对的channel，处理k线信息存储到mysql中
	// 2. 建立基于gin的http服务，从mysql中读取k线信息，然后publish到redis的指定频道中
	tradePairChan := make(chan string)
	go d.SubscribeFromRedis(config.TradePairChannelName, tradePairChan)
	// 一分钟有可能不止60条交易对哦，不能用一个固定的切片来承载，可能只能用channel了
	// 先用append来实现基本功能吧，接下来再迭代
	tradePairIn1MinList := make([]model.TradePairWithTime, 0)
	restList := make([]model.TradePairWithTime, 0)
	tradePairIn5MinList := make([]model.TradePairWithTime, 0)
	flag := 0
	var tEnd int64
	for {
		tradePairStr := <-tradePairChan
		tradePair := model.TradePairWithTime{}
		err = json.Unmarshal([]byte(tradePairStr), &tradePair)
		if err != nil {
			fmt.Printf("tradePair error: %v", err)
		}
		t := tradePair.Time
		if t%60 == 0 {
			tEnd = t + 60
		}
		// 先判断是否整数分钟，否，则先存起来；是，则清空队列后重新存起来
		// if t%60 != 0 {
		// 	tEnd = t + (60 - t%60)
		// } else {
		// 	tEnd = t + 60
		// }
		// 第一次启动的时候，判断当前时间是否为整数分钟，先将非整数分钟的数据聚合起来
		if t%60 != 0 && flag == 0 {
			restList = append(restList, tradePair)
			flag = 1
			tEnd = tradePair.Time + (60 - tradePair.Time%60)
		} else if t%60 == 0 || t%60 != 0 && t < tEnd {
			tradePairIn1MinList = append(tradePairIn1MinList, tradePair)
		}
		// 交给k线函数处理
		if tradePair.Time == tEnd {
			kLineIn1Min = kline.KLineIn1MinGen(tradePairIn1MinList)
			d.SaveKLineInfo2Mysql()
			tradePairIn1MinList = make([]model.TradePairWithTime, 0)
		}

	}
}
