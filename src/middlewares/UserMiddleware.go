package middlewares

import (
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {

}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest(ctx *gin.Context) error {
	//fmt.Println("in user middleware")
	//fmt.Println(ctx.Query("name"))
	//return errors.New("error!")
	return nil
}