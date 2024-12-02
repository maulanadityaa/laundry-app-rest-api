package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	// Add verbose logging
	log.Println("Starting database connection...")
	startTime := time.Now()

	// Collect and log environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Log environment variable lengths for debugging
	log.Printf("DB Connection Details - Host: %s, User: %s, Name: %s, Port: %s",
		maskString(dbHost),
		maskString(dbUser),
		dbName,
		dbPort)

	// Create DSN with detailed logging
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// Use context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Configure GORM with more detailed settings
	config := &gorm.Config{
		// Add more robust error handling
		Logger: logger.Default.LogMode(logger.Info),

		// Prevent N+1 queries
		PrepareStmt: true,
	}

	// Attempt connection with context
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), config)

	if err != nil {
		log.Printf("Database connection error: %v", err)
		log.Printf("Connection attempt duration: %v", time.Since(startTime))
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Get underlying SQL DB for connection pool management
	sqlDB, err := database.DB()
	if err != nil {
		log.Printf("Failed to get database connection: %v", err)
		panic(fmt.Sprintf("Failed to get database connection: %v", err))
	}

	// Aggressive connection pool management
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	// Ping database with context to verify connection
	sqlDBCtx, sqlCancel := context.WithTimeout(ctx, 10*time.Second)
	defer sqlCancel()

	if err := sqlDB.PingContext(sqlDBCtx); err != nil {
		log.Printf("Database ping error: %v", err)
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	DB = database

	log.Printf("Database connection successful. Total connection time: %v", time.Since(startTime))
}

// Helper function to mask sensitive strings
func maskString(s string) string {
	if len(s) > 4 {
		return s[:2] + strings.Repeat("*", len(s)-4) + s[len(s)-2:]
	}
	return strings.Repeat("*", len(s))
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
