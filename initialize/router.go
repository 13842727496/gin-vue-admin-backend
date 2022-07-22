package initialize

import (
	"github.com/13842727496/gin-vue-admin-backend/router"
	"github.com/gin-gonic/gin"
)

// RouterGroup general route.
type RouterGroup struct {
	loginRouter router.LoginRouter
}

// NewRouterGroup create new RouterGroup instance.
func NewRouterGroup() *RouterGroup {
	return &RouterGroup{}
}

// Routers gin routing configuration.
func (g *RouterGroup) Routers() *gin.Engine {
	r := gin.Default()
	g.loginRouter.InitLoginRouter(r)
	return r
}
