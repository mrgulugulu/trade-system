package model

type TradePair struct {
	Price float64 `gorm:"price" json:"price"`
	Amt   float64 `gorm:"amt" json:"amt"`
	Total float64 `gorm:"total" json:"total"`
}
