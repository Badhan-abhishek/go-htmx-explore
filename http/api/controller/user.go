package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"com.quizApp/http/dto"
	"com.quizApp/http/services"
	"com.quizApp/lib"
	"com.quizApp/models"

	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var input dto.CreateUserInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := services.CreateUserService(c, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Signin(c *gin.Context) {
	var input dto.SigninUserInput
	var user models.User
	err := c.ShouldBind(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":  err.Error(),
			"code": "bind_failed",
		})
		return
	}

	result := models.DB.Find(&user, "email = ?", input.Email)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":  result.Error,
			"code": "user_not_found",
		})
		return
	}

	passCheckErr := bcrypt.CompareHashAndPassword(user.Password, []byte(input.Password))
	if passCheckErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":  passCheckErr,
			"code": "pass_check_failed",
		})
		return
	}

	token, tokenGenErr := lib.GetJwt(&user)
	if tokenGenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":  tokenGenErr.Error(),
			"code": "token_gen_failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUser(c *gin.Context) {
	var users models.User
	email := c.Param("email")
	result := models.DB.Find(&users, "email = ?", email)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
