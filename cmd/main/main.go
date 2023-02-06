package main

import (
	"encoding/json"
	"trade-system/config"
	"trade-system/internal/dao"
	"trade-system/internal/kline"
	"trade-system/internal/log"
	"trade-system/internal/model"
	"trade-system/internal/server"
)

func main() {
	d, err := dao.NewDao()
	if err != nil {
		log.Sugar.Panicf("new dao error: %v", err)
		return
	}
	s := server.Server{
		Addr: config.ServiceConf.ServerCfg.Addr,
		Port: config.ServiceConf.ServerCfg.Port,
	}
	go s.Run()
	tradePairChan := make(chan string)
	go d.SubscribeFromRedis(config.TradePairChannelName, tradePairChan)
	// TODO: 用切片的话，如果遇到巨量的交易对，可能会导致消耗过多的内存空间
	tradePairIn1MinList := make([]model.TradePairWithTime, 0)
	tradePairIn5MinList := make([]model.TradePairWithTime, 0)
	flag1Min, flag5Min := 0, 0
	var tEnd1Min, tEnd5Min int64
	for {
		tradePairStr := <-tradePairChan
		tradePair := model.TradePairWithTime{}
		err = json.Unmarshal([]byte(tradePairStr), &tradePair)
		if err != nil {
			log.Sugar.Errorf("tradePair unmarshal error: %v", err)
		}
		t := tradePair.Time
		// 这里考虑到有部分数据的交易时间并不是从0秒开始，所以单独处理掉，当做是分钟线
		if t%60 != 0 && flag1Min == 0 {
			flag1Min = 1
			tEnd1Min = tradePair.Time + (60 - tradePair.Time%60)
		} else if t%60 == 0 && len(tradePairIn1MinList) == 0 {
			tEnd1Min = t + 60
		}
		if t%300 != 0 && flag5Min == 0 {
			flag5Min = 1
			tEnd5Min = tradePair.Time + (300 - tradePair.Time%300)
		} else if t%300 == 0 && len(tradePairIn5MinList) == 0 {
			tEnd5Min = t + 300
		}
		tradePairIn1MinList = append(tradePairIn1MinList, tradePair)
		tradePairIn5MinList = append(tradePairIn5MinList, tradePair)

		// 假定k线程序会一直运行，所以所有的数据都是以整分钟保存的
		switch t {
		case tEnd5Min:
			kLineIn5Min := kline.KLineIn5MinGen(tradePairIn5MinList)
			log.Sugar.Infof("kLineIn5Min: %+v", kLineIn5Min)
			err = d.SaveKLineInfo2Mysql(kLineIn5Min)
			if err != nil {
				log.Sugar.Errorf("save 5-min-k-line error: %v", err)
			}
			kByte, err := json.Marshal(kLineIn5Min)
			if err != nil {
				log.Sugar.Errorf("5-min-k-line marshal error: %v", err)
			}
			d.Publish2Redis(config.KLineIn5MinChannelName, string(kByte))
			// 重置交易对列表
			tradePairIn5MinList = make([]model.TradePairWithTime, 0)
			fallthrough
		case tEnd1Min:
			kLineIn1Min := kline.KLineIn1MinGen(tradePairIn1MinList)
			log.Sugar.Infof("kLineIn1Min: %+v", kLineIn1Min)
			err = d.SaveKLineInfo2Mysql(kLineIn1Min)
			if err != nil {
				log.Sugar.Errorf("save 1-min-k-line error: %v", err)
			}
			kByte, err := json.Marshal(kLineIn1Min)
			if err != nil {
				log.Sugar.Errorf("1-min-k-line marshal error: %v", err)
			}
			d.Publish2Redis(config.KLineIn1MinChannelName, string(kByte))
			// 重置交易对列表
			tradePairIn1MinList = make([]model.TradePairWithTime, 0)
		}
	}

	// TODO: 需要一个优雅退出
}
