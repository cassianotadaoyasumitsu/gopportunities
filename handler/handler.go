package handler

import (
	"gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	//db     *gorm.DB
	dbp *gorm.DB
)

func InitializerHandler() {
	logger = config.GetLogger("handler")
	//db = config.GETSQLite()
	dbp = config.GETPostgres()
}
