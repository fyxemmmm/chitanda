package chitanda

import (
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
	this.props = append(this.props, dba)
	return this
}


func (this *Chitanda) Earnest(group string, classes ...IClass) *Chitanda {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
		this.setProp(class)

		//reflectClass := reflect.ValueOf(class).Elem()
		//if reflectClass.NumField() > 0 {
		//	if this.dba != nil {
		//		reflectClass.Field(0).Set(reflect.New(reflectClass.Field(0).Type().Elem()))
		//		reflectClass.Field(0).Elem().Set(reflect.ValueOf(this.dba).Elem())
		//	}
		//}
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
	for i := 0; i < vClass.NumField(); i ++ {
		f := vClass.Field(i)
		if f.IsNil() == false || f.Kind() != reflect.Ptr {
			continue
		}
		if p := this.getProp(f.Type()); p != nil {
			f.Set(reflect.New(f.Type().Elem()))
			f.Elem().Set(reflect.ValueOf(p).Elem())
		}
	}
}