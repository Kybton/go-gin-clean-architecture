package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kybton/go-gin-clean-architecture/src/app/dtos"
)

// Handler for the routes which is not registered.
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Responding with the custom struct so that api consumers can handle with model
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			dtos.BaseResponse{Message: "Not Found."},
		)
	}
}
