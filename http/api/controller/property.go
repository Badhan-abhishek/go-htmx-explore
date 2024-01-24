package controller

import (
	"net/http"

	"com.backend/http/dto"
	"com.backend/lib"
	"github.com/gin-gonic/gin"
)

func AddProperty(c *gin.Context) {
	var input dto.CreatePropertyInput
	err := c.ShouldBind(&input)
	if err != nil {
		err = lib.NewHttpError("Invalid input", err.Error(), http.StatusBadRequest)
		c.Error(err)
		return
	}
}
