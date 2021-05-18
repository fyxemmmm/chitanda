package middlewares

import (
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest(ctx *gin.Context) error {
	return nil
}