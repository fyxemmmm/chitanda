package classes

import (
	"github.com/fyxemmmm/chitanda/chitanda"
	"github.com/gin-gonic/gin"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}


func (this *UserClass) UserList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "success",
		})
	}
}

func (this *UserClass) Build(chitanda *chitanda.Chitanda)  {
	chitanda.Handle("GET", "/user", this.UserList())
}
