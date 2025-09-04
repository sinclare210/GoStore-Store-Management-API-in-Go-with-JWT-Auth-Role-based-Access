package models

type User struct {
	Id       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string  `gorm:"not null" json:"name"`
	Email    string  `gorm:"uniqueIndex;not null" json:"email"`
	Password string  `gorm:"not null" json:"-"`
	Role     string  `gorm:"default:user" json:"role"`
}

