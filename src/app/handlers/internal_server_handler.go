package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/kybton/go-gin-clean-architecture/src/app/dtos"
)

func InternalServerErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		dtos.BaseResponse{Message: "Internal Server Error."},
	)
	fmt.Println(goErr)
}
