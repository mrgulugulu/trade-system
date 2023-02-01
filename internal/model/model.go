// package model 负责进行结构体的定义
package model

// TradePair 交易对的结构体
type TradePair struct {
	Price float64 `gorm:"price" json:"price"`
	Amt   float64 `gorm:"amt" json:"amt"`
	Total float64 `gorm:"total" json:"total"`
}

type TradePairWithTime struct {
	TradePair
	Time int64 `json:"time"`
}

// KLineIn1Min 1分钟k线的结构体
type KLineIn1Min struct {
	Time         int64   `gorm:"time" json:"time"`
	Open         float64 `gorm:"open" json:"open"`
	Close        float64 `gorm:"close" json:"close"`
	HighestPrice float64 `gorm:"highest_price" json:"highest_price"`
	LowestPrice  float64 `gorm:"lowest_price" json:"lowest_price"`
	Volume       float64 `gorm:"volume" json:"volume"`
}

// KLineIn5Min 5分钟k线的结构体
type KLineIn5Min struct {
	Time         int64   `gorm:"time" json:"time"`
	Open         float64 `gorm:"open" json:"open"`
	Close        float64 `gorm:"close" json:"close"`
	HighestPrice float64 `gorm:"highest_price" json:"highest_price"`
	LowestPrice  float64 `gorm:"lowest_price" json:"lowest_price"`
	Volume       float64 `gorm:"volume" json:"volume"`
}
