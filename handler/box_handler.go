package handler

import (
	"main/dto"
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

func (h *BoxHandler) CreateBox(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	var req dto.CreateBoxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, "Invalid request format: "+err.Error())
		return
	}

	box, err := h.boxService.CreateBox(userID, &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
	}

	response.CreatedResponse(c, "box successfully created", box)
}
