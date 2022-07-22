package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginApi implements methods for repository.
type LoginApi struct{}

// Login user login.
func (r *LoginApi) Login(ctx *gin.Context) {
	fmt.Println("123123")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ping",
		"code":    0,
		"msg":     "登录成功",
	})
}
