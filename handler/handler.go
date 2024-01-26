package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopportunities/config"
	"gorm.io/gorm"
	"net/http"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

var (
	Logger *config.Logger
	DB     *gorm.DB
)

func InitializerHandler() {
	Logger = config.GetLogger("handler")
	DB = config.GETSQLite()
}

func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s was successful", op),
		"data":    data,
	})
}

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type%s) is required", name, typ)
}
