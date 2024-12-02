package controllers

import (
	"fmt"

	"github.com/eDyrr/learning-gorm/models"
	"gorm.io/gorm"
)

func AddUser(db *gorm.DB, user *models.User) {
	result := db.Create(&user)
	if result.Error != nil {
		panic("couldnt create a new user")
	}
}

func SearchByFirstName(db *gorm.DB, firstName string) {
	var users []models.User
	result := db.Where("FirstName = ?", firstName).Find(&users)
	if result.Error != nil {
		panic("failed to retrieve user")
	}
}

func ListAll(db *gorm.DB) {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		panic("failed to retrieve users")
	}

	for _, user := range users {
		fmt.Printf("User ID: %d, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
	}
}

func UpdateFirstName(db *gorm.DB, id int, firstName string) {
	var user models.User
	result := db.First(&user, id)

	if result.Error != nil {
		panic("failed to find user")
	}

	user.FirstName = firstName

	result = db.Save(&user)

	if result.Error != nil {
		panic("failed to update user")
	}

	fmt.Println("user update success")
}
