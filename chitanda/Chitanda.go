package chitanda

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type Chitanda struct {
	*gin.Engine
	g *gin.RouterGroup
	dba interface{}
}

func Inquisitive() *Chitanda {
	ctd :=  &Chitanda{Engine: gin.New()}
	ctd.Use(ErrorHandler())
	return ctd
}

func (this *Chitanda) Start() {
	this.Run(":8080")
}

func (this *Chitanda) Handle(httpMethod, relativePath string, handler interface{}) *Chitanda {
	if h:= Convert(handler);h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}

	return this
}

func (this *Chitanda) Responsible(f Responsible) *Chitanda{
	this.Use(func(context *gin.Context) {
		err := f.OnRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		}else {
			context.Next()
		}
	})
	return this
}

func (this *Chitanda) Joyful(dba interface{}) *Chitanda {
	this.dba = dba
	return this
}


func (this *Chitanda) Earnest(group string, classes ...IClass) *Chitanda {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)  // class是控制器类
		reflectClass := reflect.ValueOf(class).Elem()
		if reflectClass.NumField() > 0 {
			if this.dba != nil {
				reflectClass.Field(0).Set(reflect.New(reflectClass.Field(0).Type().Elem()))
				reflectClass.Field(0).Elem().Set(reflect.ValueOf(this.dba).Elem())
			}
		}
	}
	return this
}

