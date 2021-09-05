package model

type User struct {
	UserId   uint   `gorm:"primaryKey;not null"`
	UserName string `gorm:"not null;default:null"`
	Email    string `gorm:"not null;default:null;unique"`
	Password string `gorm:"not null;default:null"`
}
