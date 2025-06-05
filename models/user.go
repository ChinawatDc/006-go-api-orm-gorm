package models

import "gorm.io/gorm"

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (User, error) {
	var user User
	result := db.First(&user, id)
	return user, result.Error
}

func CreateUser(user User) error {
	result := db.Create(&user)
	return result.Error
}
