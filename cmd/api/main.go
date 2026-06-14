package main

import (
	"github.com/fleece30/wdmmg/config"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *gorm.DB
	config config.Conf
	router *gin.Engine
}

func InitServer(db *gorm.DB, config config.Conf) (*Server, error) {
	server := &Server{
		db:     db,
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// func main() {
// 	config, err := config
// }
