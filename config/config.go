package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	//db     *gorm.DB
	dbp    *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize DB
	//db, err = InitializeSQLite()
	//if err != nil {
	//	return fmt.Errorf("error initializing SQLite: %v", err)
	//}

	dbp, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("error initializing Postgre: %v", err)
	}

	return nil
}

//	func GETSQLite() *gorm.DB {
//		return db
//	}
func GETPostgres() *gorm.DB {
	return dbp
}

func GetLogger(p string) *Logger {
	// Initialize the logger
	logger = NewLogger(p)
	return logger
}
