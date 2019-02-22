package web

import (
	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/go-clean/src/adapter/controllers"
	"github.com/yuki-toida/go-clean/src/registry/interfaces"
)

func NewRouter(u interfaces.UseCase) *gin.Engine {
	router := gin.Default()

	uu := u.NewUserUseCase()
	uc := controllers.NewUserController(uu)

	users := router.Group("/users")
	{
		users.GET("/", uc.Find)
		users.POST("/", uc.Create)
		users.GET("/:uid", uc.First)
		users.DELETE("/:uid", uc.Delete)
	}

	eu := u.NewEmailUseCase()
	ec := controllers.NewEmailController(eu)

	emails := router.Group("/emails")
	{
		emails.POST("/", ec.Create)
		emails.PATCH("/", ec.Update)
	}

	return router
}
