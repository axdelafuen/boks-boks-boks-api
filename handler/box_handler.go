package handler

import (
	"main/middleware"
	"main/response"
	"main/service"

	"github.com/gin-gonic/gin"
)

type BoxHandler struct {
	boxService *service.BoxService
}

func NewBoxHandler(boxService *service.BoxService) *BoxHandler {
	return &BoxHandler{
		boxService: boxService,
	}
}

func (h *BoxHandler) GetBoxes(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	boxes, err := h.boxService.GetBoxes(userID.String())
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessResponse(c, 200, "Boxes fetched", boxes)
}
