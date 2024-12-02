package config

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	// err = database.AutoMigrate(&entity.Role{}, &entity.Account{}, &entity.Customer{}, &entity.Employee{}, &entity.Product{}, &entity.Transaction{}, &entity.TransactionDetail{})
	// if err != nil {
	// 	panic("Failed to migrate database")
	// }

	sqlDB, err := database.DB()
	if err != nil {
		panic("Failed to get database connection")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = database

	// initRole()
}

// func initRole() {
// 	roles := []entity.Role{
// 		{
// 			ID:   uuid.NewString(),
// 			Name: "ROLE_EMPLOYEE",
// 		},
// 		{
// 			ID:   uuid.NewString(),
// 			Name: "ROLE_CUSTOMER",
// 		},
// 	}

// 	for _, role := range roles {
// 		var roleExist entity.Role

// 		result := DB.Where("name = ?", role.Name).First(&roleExist).Debug()
// 		if result.Error != nil {
// 			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 				DB.Create(&role)
// 			} else {
// 				fmt.Println(result.Error)
// 			}
// 		} else {
// 			fmt.Println("Role already exist")
// 		}
// 	}
// }
