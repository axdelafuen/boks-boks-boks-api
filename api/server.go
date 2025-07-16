package api

import (
  "os"
  "log"
  "net/http"

  "github.com/gin-gonic/gin"

  "main/database"
)

type Server struct {
  router *gin.Engine
  dbManager *database.Manager
  jwtSecret string
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

  server := &Server {
    router: gin.Default(),
    dbManager: dbManager,
    jwtSecret: jwtSecret,
  }

  server.setupRoutes()
  return server, nil
}

func (s *Server) setupRoutes() {
  s.router.GET("/hello", func(c *gin.Context) {
    data := gin.H{
      "message": "Hello World!",
      "status": "success",
    }

    c.JSON(http.StatusOK, data)
  })
}

func (s *Server) Run(addr ...string) error {
  return s.router.Run(addr...)
}

func (s *Server) Close() error {
  return s.dbManager.Close()
}
