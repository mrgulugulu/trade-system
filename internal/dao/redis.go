package dao

import "fmt"

// SubscribeFromRedis 从redis中订阅指定channel
func (d *dao) SubscribeFromRedis(chanName string, tradePairChan chan<- string) {
	pubSub := d.RedisDb.Subscribe(chanName)
	defer pubSub.Close()
	for msg := range pubSub.Channel() {
		tradePairChan <- msg.Payload
	}
}

// Publish2Redis publish信息到redis的指定channel
func (d *dao) Publish2Redis(chanName string, msg string) {
	_, err := d.RedisDb.Publish(chanName, msg).Result()
	if err != nil {
		fmt.Printf("data: %s publish to redis error: %v", msg, err)
		// TODO: 之后要保存到日志中
	}
}
