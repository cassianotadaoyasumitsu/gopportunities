package config

import (
	"gopportunities/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Create DB
	db, err := gorm.Open(postgres.Open("host=localhost user=seu_user_name password=seu_password dbname=seu_dbname port=5432 sslmode=disable TimeZone=Asia/Tokyo"), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error Postgres opening database: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error Postgres migrating database: %v", err)
		return nil, err
	}

	return db, nil
}

//func InitializeSQLite() (*gorm.DB, error) {
//	logger := GetLogger("sqlite")
//
//	// Check if DB exists
//	dbPath := "./db/main.db"
//	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
//		logger.Infof("SQLite database does not exist, creating...")
//		// Create DB
//		err := os.MkdirAll("./db", os.ModePerm)
//		if err != nil {
//			logger.Errorf("Error SQLite creating database directory: %v", err)
//			return nil, err
//		}
//		file, err := os.Create(dbPath)
//		if err != nil {
//			logger.Errorf("Error SQLite creating database file: %v", err)
//			return nil, err
//		}
//		err = file.Close()
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	// Create DB
//	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
//	if err != nil {
//		logger.Errorf("Error SQLite opening database: %v", err)
//		return nil, err
//	}
//	// Migrate the schema
//	err = db.AutoMigrate(&schemas.Opening{})
//	if err != nil {
//		logger.Errorf("Error SQLite migrating database: %v", err)
//		return nil, err
//	}
//
//	return db, nil
//}
