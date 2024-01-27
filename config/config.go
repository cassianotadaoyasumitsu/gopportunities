package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	dbp    *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	dbp, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("error initializing Postgre: %v", err)
	}

	return nil
}

func GETPostgres() *gorm.DB {
	return dbp
}

func GetLogger(p string) *Logger {
	// Initialize the logger
	logger = NewLogger(p)
	return logger
}
