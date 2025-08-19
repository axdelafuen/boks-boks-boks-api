package handler

import (
  "main/service"
  "main/middleware"
  "main/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
  userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
  return &UserHandler{
    userService: userService,
  }
}

func (h *UserHandler) GetUser(c *gin.Context) {
  userId, err := middleware.GetUserIDFromContext(c)
  if err != nil {
    response.InternalServerError(c, err.Error())
    return
  }

  username := c.Param("username")

  user, err := h.userService.GetUser(userId.String(), username)
  if err != nil {
    response.InternalServerError(c, err.Error())
    return 
  }

  response.OKResponse(c, "User datas fetched", user)
}
