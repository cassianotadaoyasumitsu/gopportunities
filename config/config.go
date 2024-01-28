package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize DB
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing SQLite: %v", err)
	}

	return nil
}

func GETSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize the logger
	logger = NewLogger(p)
	return logger
}

func LoadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading env variables: %v", err)
	}
	return nil
}
