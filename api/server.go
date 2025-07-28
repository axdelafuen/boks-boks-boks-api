package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"main/database"
	"main/handler"
	"main/middleware"
	"main/service"
)

type Server struct {
	router       *gin.Engine
	dbManager    *database.Manager
	jwtSecret    string
	authHandler  *handler.AuthHandler
	boxHandler   *handler.BoxHandler
	itemHandler  *handler.ItemHandler
	labelHandler *handler.LabelHandler
}

func NewServer() (*Server, error) {
	dbManager, err := database.NewManager()
	if err != nil {
		return nil, err
	}

	jwtSecret := os.Getenv("SECRET_TOKEN")
	if jwtSecret == "" {
		log.Fatal("SECRET_TOKEN env variable is required")
	}

	authService := service.NewAuthService(dbManager.GetDB(), jwtSecret)
	boxService := service.NewBoxService(dbManager.GetDB())
	itemService := service.NewItemService(dbManager.GetDB())
	labelService := service.NewLabelService(dbManager.GetDB())

	authHandler := handler.NewAuthHandler(authService)
	boxHandler := handler.NewBoxHandler(boxService)
	itemHandler := handler.NewItemHandler(itemService)
	labelHandler := handler.NewLabelHandler(labelService)

	server := &Server{
		router:       gin.Default(),
		dbManager:    dbManager,
		jwtSecret:    jwtSecret,
		authHandler:  authHandler,
		boxHandler:   boxHandler,
		itemHandler:  itemHandler,
		labelHandler: labelHandler,
	}

	// Configure CORS middleware
	server.router.Use(middleware.CORSMiddleware())

	server.setupRoutes()
	return server, nil
}

func (s *Server) setupRoutes() {
	s.router.GET("/hello", func(c *gin.Context) {
		data := gin.H{
			"message": "Hello World!",
			"status":  "success",
		}
		c.JSON(http.StatusOK, data)
	})
	s.router.POST("/login", s.authHandler.Login)
	s.router.POST("/register", s.authHandler.Register)

	api := s.router.Group("/api")
	api.Use(middleware.AuthMiddleware(s.jwtSecret))
	{
		api.GET("/boxes", s.boxHandler.GetBoxes)
		api.POST("/boxes", s.boxHandler.CreateBox)
		api.PUT("/boxes", s.boxHandler.UpdateBox)
		api.DELETE("/boxes/:id", s.boxHandler.DeleteBox)

		api.GET("/boxes/:id/items", s.itemHandler.GetItems)
		api.POST("/boxes/:id/items", s.itemHandler.CreateItem)
		api.DELETE("/boxes/:id/items/:itemid", s.itemHandler.DeleteItem)
		api.PUT("/boxes/:id/items", s.itemHandler.UpdateItem)

		api.POST("/items/:itemid/labels/:labelid", s.labelHandler.AddLabelToItem)
		api.GET("/labels", s.labelHandler.GetLabel)
		api.POST("/labels", s.labelHandler.CreateLabel)
		api.DELETE("/labels/:id", s.labelHandler.DeleteLabel)
	}
}

func (s *Server) Run(addr ...string) error {
	return s.router.Run(addr...)
}

func (s *Server) Close() error {
	return s.dbManager.Close()
}
