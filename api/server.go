package api

import (
  "os"
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

type Server struct {
  router *gin.Engine
  jwtSecret string
}

func NewServer() (*Server, error) {
  jwtSecret := os.Getenv("SECRET_TOKEN")
  if jwtSecret == "" {
    log.Fatal("SECRET_TOKEN env variable is required")
  }

  server := &Server {
    router: gin.Default(),
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
  // Close the Database with this
  return nil
}
