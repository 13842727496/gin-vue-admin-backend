package repository

import "github.com/gin-gonic/gin"

// LoginRepository interface of Login storage.
type LoginRepository interface {
	Login(ctx *gin.Context)
}
