package model

type Member struct {
	Id           int64   `gorm:"AUTO_INCREMENT" json:"id"`
	UserName     string  `json:"user_name"`
	Mobile       string  `json:"mobile"`
	Password     string  `json:"password"`
	RegisterTime int64   `json:"register_time"`
	Avatar       string  `json:"avatar"`
	Balance      float64 `json:"balance"`
	IsActive     int8    `json:"is_active"`
	City         string  `json:"city"`
}
