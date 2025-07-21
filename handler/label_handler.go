package handler

import (
	"main/dto"
	"main/middleware"
	"main/response"
	"main/service"

	"github.com/gin-gonic/gin"
)

type LabelHandler struct {
	labelService *service.LabelService
}

func NewLabelHandler(labelService *service.LabelService) *LabelHandler {
	return &LabelHandler{
		labelService: labelService,
	}
}

func (h *LabelHandler) CreateLabel(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	var req dto.CreateLabelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	label, err := h.labelService.CreateLabel(userID.String(), &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessResponse(c, 201, "label created", label)
}

func (h *LabelHandler) GetLabel(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	labels, err := h.labelService.GetLabels(userID.String())
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessResponse(c, 201, "Labels fetched", labels)
}
