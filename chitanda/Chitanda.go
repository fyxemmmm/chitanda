package chitanda

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Chitanda struct {
	*gin.Engine
	g *gin.RouterGroup
	props []interface{}
}

func Inquisitive() *Chitanda {
	ctd :=  &Chitanda{Engine: gin.New(), props: make([]interface{}, 0)}
	ctd.Use(ErrorHandler())
	return ctd
}

func (this *Chitanda) Start() {
	config := InitConfig()
	this.Run(fmt.Sprintf(":%d", config.Server.Port))
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

func (this *Chitanda) Joyful(beans ...interface{}) *Chitanda {
	this.props = append(this.props, beans)
	return this
}


func (this *Chitanda) Earnest(group string, classes ...IClass) *Chitanda {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
		this.setProp(class)

	}
	return this
}

func (this *Chitanda) getProp(t reflect.Type) interface{} {
	for _, p := range this.props {
		if t == reflect.TypeOf(p) {
			return p
		}
	}
	return nil
}

func (this *Chitanda) setProp(class IClass) {
	vClass := reflect.ValueOf(class).Elem()
	vClassT := reflect.TypeOf(class).Elem()
	for i := 0; i < vClass.NumField(); i ++ {
		f := vClass.Field(i)
		if f.IsNil() == false || f.Kind() != reflect.Ptr {
			continue
		}
		if p := this.getProp(f.Type()); p != nil {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
			if IsAnnotation(f.Type()) {
				p.(Annotation).SetTag(vClassT.Field(i).Tag)
			}

		}
	}
}