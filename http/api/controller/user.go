package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"com.backend/http/dto"
	"com.backend/http/services"
	"com.backend/lib"
)

func Signup(c *gin.Context) {
	var input dto.CreateUserInput
	if err := c.ShouldBind(&input); err != nil {
		err = lib.NewHttpError("Field validation failed", err.Error(), http.StatusBadRequest)
		c.Error(err)
		return
	}
	user, err := services.CreateUserService(input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &user,
	})
}

func Signin(c *gin.Context) {
	var input dto.SigninUserInput
	err := c.ShouldBind(&input)
	if err != nil {
		err = lib.NewHttpError("Field validation failed", err.Error(), http.StatusBadRequest)
		c.Error(err)
		return
	}
	token, err := services.SignInUserService(input)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsersService()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUser(c *gin.Context) {
	user, err := services.GetUserService(c.Param("email"))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func GetCurrentUser(c *gin.Context) {
	currentEmail, err := lib.GetAuthEmail(c)
	if err != nil {
		c.Error(err)
		return
	}

	user, err := services.GetUserWithLandlordService(*currentEmail)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	var input dto.UpdateUserInput

	if err := c.ShouldBind(&input); err != nil {
		err = lib.NewHttpError("Field validation failed", err.Error(), http.StatusBadRequest)
		c.Error(err)
		return
	}

	currentEmail, err := lib.GetAuthEmail(c)
	if err != nil {
		c.Error(err)
		return
	}

	user, err := services.UpdateUserService(currentEmail, input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
