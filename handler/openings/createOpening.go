package openings

import (
	"github.com/gin-gonic/gin"
	"gopportunities/handler"
	"gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1

// CreateOpeningHandler
// @Summary Create a new opening
// @Description Create a new job opening
// @Tags opening
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	user, userID := handler.GetUserInfo(ctx)

	request := CreateOpeningRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("Error while validating request: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
		UserID:   userID,
		User:     user, // TODO: check if this is necessary
	}

	if err := handler.DB.Create(&opening).Error; err != nil {
		handler.Logger.Errorf("Error while creating a new opening: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "Error while creating a new opening")
		return
	}

	handler.SendSuccess(ctx, "CreateOpeningHandler", opening)
}
