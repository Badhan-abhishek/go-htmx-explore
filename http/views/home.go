package views

import (
	"com.quizApp/http/dto"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func Swap(c *gin.Context) {
	c.HTML(200, "swap.html", gin.H{})
}

func SignIn(c *gin.Context) {
	var input dto.SigninUserInput
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(200, "index.html", gin.H{
			"err":  err.Error(),
			"code": "validation_failed",
		})
		return
	}
}
