package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"main/database"
	"main/handler"
	"main/service"
)

type Server struct {
	router      *gin.Engine
	dbManager   *database.Manager
	jwtSecret   string
	authHandler *handler.AuthHandler
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

	authHandler := handler.NewAuthHandler(authService)

	server := &Server{
		router:      gin.Default(),
		dbManager:   dbManager,
		jwtSecret:   jwtSecret,
		authHandler: authHandler,
	}

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
}

func (s *Server) Run(addr ...string) error {
	return s.router.Run(addr...)
}

func (s *Server) Close() error {
	return s.dbManager.Close()
}
