package models

type Order struct {
	ID           uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint    `gorm:"not null" json:"user_id"`
	ProductID    uint    `gorm:"not null" json:"product_id"`
	ProductName  string  `gorm:"not null" json:"product_name"`
	ProductPrice float64 `json:"product_price"`	
}

