package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/go-clean/src/application/usecase"
)

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(u usecase.UserUseCase) *userController {
	return &userController{userUseCase: u}
}

func (uc *userController) Find(c *gin.Context) {
	users, err := uc.userUseCase.Find()
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, users)
}

func (uc *userController) First(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("uid"))
	user, err := uc.userUseCase.First(uint64(id))
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func (uc *userController) Create(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	user, err := uc.userUseCase.Create(name)
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func (uc *userController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("uid"))
	if err := uc.userUseCase.Delete(uint64(id)); err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.JSON(http.StatusOK, "DELETED")
}
