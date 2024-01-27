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
