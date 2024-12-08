package main

import (
	"github.com/eDyrr/learning-gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:edd7355608@tcp(127.0.0.1:3306)/gormDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to DB")
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		panic("failed to perform migration:" + err.Error())
	}

	// newUser := models.User{
	// 	FirstName: "Jane",
	// 	LastName:  "Doe",
	// 	Email:     "janedoe@gmail.com",
	// 	Country:   "Spain",
	// 	Role:      "Chef",
	// 	Age:       30,
	// }

	// result := db.Create(&newUser)
	// if result.Error != nil {
	// 	panic("failed to create user:" + result.Error.Error())
	// }

	SearchByFirstName(db, "Jane")
}
