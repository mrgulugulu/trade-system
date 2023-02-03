package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	// 先用append来实现基本功能吧，接下来再迭代
	tradePairIn1MinList := make([]model.TradePairWithTime, 0)
	// tradePairIn5MinList := make([]model.TradePairWithTime, 0)
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
		// 这里考虑到有部分数据的交易时间并不是从0秒开始，所以单独处理掉，当做是分钟线
		if t%60 != 0 && flag == 0 {
			flag = 1
			tEnd = tradePair.Time + (60 - tradePair.Time%60)
		} else if t%60 == 0 && len(tradePairIn1MinList) == 0 {
			tEnd = t + 60
		}
		tradePairIn1MinList = append(tradePairIn1MinList, tradePair)

		// 假定k线程序会一直运行，所以所有的数据都是以整分钟保存的
		if tradePair.Time == tEnd {
			kLineIn1Min := kline.KLineIn1MinGen(tradePairIn1MinList)
			fmt.Printf("%v", kLineIn1Min)
			// err = d.SaveKLineInfo2Mysql(kLineIn1Min)
			if err != nil {
				log.Printf("save k line error: %v", err)
			}
			kByte, err := json.Marshal(kLineIn1Min)
			if err != nil {
				log.Printf("marshal error: %v", err)
			}
			d.Publish2Redis(config.KLineIn1MinChannelName, string(kByte))
			// 重置
			tradePairIn1MinList = make([]model.TradePairWithTime, 0)
		}

	}
}
