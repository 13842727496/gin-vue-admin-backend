package main

import (
	"github.com/13842727496/gin-vue-admin-backend/initialize"
)

func main() {
	r := initialize.NewRouterGroup().Routers()
	r.Run(":8888")
}
