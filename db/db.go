package db

import (
	"fmt"

	"github.com/sharmarajdaksh/go-pwd/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Initialize initializes the database connection for password access
func Initialize() error {

	dbFile := config.GetDatabaseFile()

	var err error
	db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = db.AutoMigrate(&Password{}); err != nil {
		return fmt.Errorf("failed to make database migrations: %w", err)
	}

	return nil
}
