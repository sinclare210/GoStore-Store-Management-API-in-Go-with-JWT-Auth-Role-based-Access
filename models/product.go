package models


type Product struct {
	Id          uint `gorm:"primaryKey;autoIncrement"`
	Name        string  `binding:"required"`
	Description string  `binding:"required"`
	Price       float64 `binding:"required"`
	Quantity    int64   `binding:"required"`
	UserID     	uint 
}

