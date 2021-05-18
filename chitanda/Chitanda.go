package chitanda

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Chitanda struct {
	*gin.Engine
	g *gin.RouterGroup
	//props []interface{}
	beanFactory *BeanFactory
}

func Inquisitive() *Chitanda {
	ctd :=  &Chitanda{Engine: gin.New(), beanFactory:NewBeanFactory()}
	ctd.Use(ErrorHandler())
	ctd.beanFactory.setBean(InitConfig())  //整个配置加载进bean中
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
	this.beanFactory.setBean(beans...)
	return this
}


func (this *Chitanda) Earnest(group string, classes ...IClass) *Chitanda {
	this.g=this.Group(group)
	for _,class:=range classes{
		class.Build(this)  //这一步是关键 。 这样在main里面 就不需要 调用了
		this.beanFactory.inject(class)
	}
	return this
}

