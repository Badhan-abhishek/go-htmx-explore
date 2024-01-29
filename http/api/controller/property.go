package controller

import (
	"net/http"

	"com.backend/http/dto"
	"com.backend/http/services"
	"com.backend/lib"
	"github.com/gin-gonic/gin"
)

func AddProperty(c *gin.Context) {
	var input dto.CreatePropertyInput
	if err := c.ShouldBind(&input); err != nil {
		err = lib.NewHttpError("Field validation failed", err.Error(), http.StatusBadRequest)
		c.Error(err)
		return
	}

	landlord, err := lib.GetAuthLandlord(c)

	if err != nil {
		c.Error(err)
		return
	}

	property, err := services.AddPropertyService(input, landlord)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": property,
	})
}

func GetLandlordProperties(c *gin.Context) {
	landlord, err := lib.GetAuthLandlord(c)
	if err != nil {
		c.Error(err)
		return
	}

	properties, err := services.GetLandlordPropertiesService(landlord)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": properties,
	})
}

func CreateUnit(c *gin.Context) {
	var input dto.UnitInput
	if err := c.ShouldBind(&input); err != nil {
		err = lib.NewHttpError("Field validation failed", err.Error(), http.StatusBadRequest)
		c.Error(err)
		return
	}

	_, err := lib.GetAuthLandlord(c)
	if err != nil {
		c.Error(err)
		return
	}

	unit, err := services.AddUnitToPropertyService(input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": unit,
	})

}
