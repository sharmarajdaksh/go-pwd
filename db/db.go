package db

import (
	"fmt"

	"github.com/sharmarajdaksh/go-pwd/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitializeDB initializes the database connection for password access
func InitializeDB() error {

	dbFile := config.GetDatabaseFile()

	var err error
	db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	if err = db.AutoMigrate(&Password{}); err != nil {
		return fmt.Errorf("failed to make database migrations: %v", err)
	}

	return nil
}
