package router

import (
	"github.com/13842727496/gin-vue-admin-backend/initialize"
	"github.com/gin-gonic/gin"
)

// LoginRouter assign routes.
type LoginRouter struct{}

// InitLoginRouter initialize the login route.
func (l *LoginRouter) InitLoginRouter(router *gin.Engine) {
	loginRouter := router.Group("base")
	{
		loginRouter.GET("login", initialize.NewApiGroup().LoginRepository.Login)
		loginRouter.POST("login", initialize.NewApiGroup().LoginRepository.Login)
	}
}
