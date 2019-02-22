package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/go-clean/src/application/usecase"
)

type emailController struct {
	emailUseCase usecase.EmailUseCase
}

func NewEmailController(u usecase.EmailUseCase) *emailController {
	return &emailController{emailUseCase: u}
}

func (ec *emailController) Create(c *gin.Context) {
	emailID, _ := strconv.Atoi(c.PostForm("eid"))
	email := c.PostForm("email")
	e, err := ec.emailUseCase.Create(uint64(emailID), email)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, e)
}

func (ec *emailController) Update(c *gin.Context) {
	emailID, _ := strconv.Atoi(c.PostForm("eid"))
	email := c.PostForm("email")
	e, err := ec.emailUseCase.Update(uint64(emailID), email)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, e)
}
