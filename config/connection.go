package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	// Extremely lightweight connection with minimal resources
	log.Println("Initializing database connection...")
	startTime := time.Now()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Ultra-conservative GORM configuration
	config := &gorm.Config{
		// Disable unnecessary features
		SkipDefaultTransaction: true,

		// Minimal logging
		Logger: logger.Default.LogMode(logger.Silent),

		// Prevent N+1 query issues
		PrepareStmt: false,
	}

	// Use context with very short timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), config)

	if err != nil {
		log.Printf("Database connection error: %v", err)
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Get underlying SQL DB with minimal connection pool
	sqlDB, err := database.DB()
	if err != nil {
		log.Printf("Failed to get database connection: %v", err)
		panic(fmt.Sprintf("Failed to get database connection: %v", err))
	}

	// Ultra-conservative connection pool for Hobby plan
	sqlDB.SetMaxIdleConns(2)                  // Minimum idle connections
	sqlDB.SetMaxOpenConns(5)                  // Low max open connections
	sqlDB.SetConnMaxLifetime(3 * time.Minute) // Shorter connection lifetime

	// Quick ping with very short timeout
	pingCtx, pingCancel := context.WithTimeout(ctx, 5*time.Second)
	defer pingCancel()

	if err := sqlDB.PingContext(pingCtx); err != nil {
		log.Printf("Database ping error: %v", err)
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	DB = database

	log.Printf("Database connection completed in %v", time.Since(startTime))
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
