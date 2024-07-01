package app

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kybton/go-gin-clean-architecture/configs"
	"github.com/kybton/go-gin-clean-architecture/src/app/handlers"
	"github.com/kybton/go-gin-clean-architecture/src/app/middlewares"
	"github.com/kybton/go-gin-clean-architecture/src/app/routers"
	"go.uber.org/dig"
)

// Struct for the REST Api.
// Changes of the framework has to be make here.
type Server struct {
	router *gin.Engine

	// Define every routes for the router engine here
	RootRouter routers.RootRouter
}

// Struct to hold the dependencies for the server which will later be used to inject.
type ServerDeps struct {
	dig.In

	// Dependencies for the server
	RootRouter routers.RootRouter `name:"RootRouter"`
}

func NewServer(deps ServerDeps) {
	switch configs.Configurations.Server.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
		os.Exit(4)
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	server := Server{
		router:     gin.Default(),
		RootRouter: deps.RootRouter,
	}

	server.router.ForwardedByClientIP = true
	server.router.SetTrustedProxies([]string{"127.0.0.1"})

	server.run()
}

// Method to register all the routes to the router.
func (s *Server) attatchRoutes() {
	rootIndex := s.router.Group("/")
	s.RootRouter.AddRoutes(rootIndex)
}

// Method that holds the steps which is necessary to run the server.
func (s *Server) run() {
	// Registering CORS middleware
	s.router.Use(middlewares.CorsMiddleware())

	// Registering custom recovery handler for custom response
	s.router.Use(gin.CustomRecovery(handlers.InternalServerErrorHandler))

	// Registering no route handler for custom response
	s.router.NoRoute(handlers.NoRouteHandler())

	s.attatchRoutes()

	s.router.Run(fmt.Sprintf("%s:%s",
		configs.Configurations.Server.Host,
		configs.Configurations.Server.Port,
	))
}
