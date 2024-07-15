package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kybton/go-gin-clean-architecture/src/app/dtos"
)

type RootController struct{}

func NewRootController() RootController {
	return RootController{}
}

func (ctrl *RootController) Root(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		dtos.BaseResponse{
			Title:   "Success",
			Message: "Server is up.",
		},
	)
}
