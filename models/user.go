package models

type User struct {
	Userid      string `gorm:"primaryKey"` // This is the primary key
	PhoneNumber string
	Username    string
	UciID       string
}
