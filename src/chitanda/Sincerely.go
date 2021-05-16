package chitanda

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

var SincerelyList []Sincerely

type Sincerely interface {
	SincerelyTo() gin.HandlerFunc
}

func init() {
	SincerelyList = []Sincerely{ //返回这个类型的指针
		new(SincerelyString),
		new(SincerelyModel),
		new(SincerelyModels),
	}
}

// func(ctx *gin.Context) string 断言失败 SincerelyString
// func(ctx *gin.Context) string 断言失败 Sincerely

func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler) // func(ctx *gin.Context) string

	for _, r := range SincerelyList { // 类型指针
		r_ref := reflect.ValueOf(r).Elem()
		if h_ref.Type().ConvertibleTo(r_ref.Type()) {
			r_ref.Set(h_ref) // 反射设置值
			return r_ref.Interface().(Sincerely).SincerelyTo()
		}
	}
	return nil
}

type SincerelyModel func(ctx *gin.Context) Model

func (this SincerelyModel) SincerelyTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, this(context))
	}
}

type SincerelyModels func(ctx *gin.Context) Models

func (this SincerelyModels) SincerelyTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "application/json")
		context.Writer.WriteString(string(this(context)))
	}
}

type SincerelyString func(ctx *gin.Context) string

func (this SincerelyString) SincerelyTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, this(context))
	}
}

