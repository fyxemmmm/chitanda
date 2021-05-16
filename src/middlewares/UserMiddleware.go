package middlewares

import (
	"fmt"
)

type UserMiddleware struct {

}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest() error {
	fmt.Println("in user middleware")
	//return errors.New("error!")
	return nil
}