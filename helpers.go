package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func reqValidator(c *gin.Context) (*CarRequest, error) {
	model := c.PostForm("model")
	if model == "" {
		return nil, errors.New("model invalid")
	}

	color := c.PostForm("color")
	if color == "" {
		return nil, errors.New("color invalid")
	}

	brand := c.PostForm("brand")
	if brand == "" {
		return nil, errors.New("brand invalid")
	}

	newReq := CarRequest{
		Model: model,
		Color: color,
		Brand: brand,
	}

	return &newReq, nil
}
