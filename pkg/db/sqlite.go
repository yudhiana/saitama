package db

import (
	"log"
	"sync"

	"saitama/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbOnce sync.Once
var dbConnection *gorm.DB

// GetDatabaseConnection returns the database connection with a singleton pattern
func GetDatabaseConnection() *gorm.DB {
	dbOnce.Do(func() {
		// connect to database
		// This will automatically create 'users.db' if it doesn't exist
		db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database:", err)
		}

		dbConnection = db.Debug()

		// Create the table (if not exists)
		err = dbConnection.AutoMigrate(&models.User{})
		if err != nil {
			log.Fatal("failed to migrate schema:", err)
		}

		log.Println("SQLite database and users table created successfully!")
	})
	return dbConnection
}
