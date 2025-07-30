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
		return
	}

	response.CreatedResponse(c, "box successfully created", box)
}

func (h *BoxHandler) DeleteBox(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	boxID := c.Param("id")

	if err := h.boxService.DeleteBox(userID, boxID); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.OKResponse(c, "box and related items deleted", nil)
}

func (h *BoxHandler) UpdateBox(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	var req dto.UpdateBoxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, "invalid request format: "+err.Error())
		return
	}

	if err := h.boxService.UpdateBox(userID, req); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.OKResponse(c, "box updated", nil)
}

func (h *BoxHandler) GetBoxContainItemWithTitle(c *gin.Context) {
	userId, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	itemTitle := c.Param("title")

	boxId, err := h.boxService.GetBoxContainItemWithTitle(userId, itemTitle)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.OKResponse(c, "box id find", boxId)
}
