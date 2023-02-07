package dao

import (
	"time"
	"trade-system/internal/log"
)

// SubscribeFromRedis 从redis中订阅指定channel
func (d *dao) SubscribeFromRedis(chanName string, tradePairChan chan<- string) {
	pubSub := d.RedisDb.Subscribe(chanName)
	defer pubSub.Close()
	for msg := range pubSub.Channel() {
		tradePairChan <- msg.Payload
	}
}

// Publish2Redis publish信息到redis的指定channel
func (d *dao) Publish2Redis(chanName string, msgChan <-chan string) {
	for v := range msgChan {
		_, err := d.RedisDb.Publish(chanName, v).Result()
		if err != nil {
			log.Sugar.Errorf("data: %s publish to redis error: %v", v, err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
