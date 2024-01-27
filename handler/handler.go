package handler

import (
	"gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	dbp    *gorm.DB
)

func InitializerHandler() {
	logger = config.GetLogger("handler")
	dbp = config.GETPostgres()
}
