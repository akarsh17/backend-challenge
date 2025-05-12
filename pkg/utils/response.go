package utils

import (
	"backend-challenge/pkg/errors"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, SuccessResponse{Data: data})
}

func Error(c *gin.Context, err errors.APIError) {
	c.JSON(err.Code, err)
}
