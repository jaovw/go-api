package config

import (
	"os"

	"github.com/jaovw/go-api/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// [joaovictor - 21/05/2025] Check if the DB file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating ...")
		// [joaovictor - 21/05/2025] Create the DB file and directory
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// [joaovictor - 21/05/2025] Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("SQLite opening error: %v", err)
		return nil, err
	}

	// [joaovictor - 21/05/2025] Migrate the Schema
	err = db.AutoMigrate(&schema.Opening{})

	if err != nil {
		logger.Errorf("SQLite autoMigration error: %v", err)
		return nil, err
	}

	return db, err
}
