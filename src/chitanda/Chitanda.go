package chitanda

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Chitanda struct {
	*gin.Engine
	g *gin.RouterGroup
	beanFactory *BeanFactory
}

func Inquisitive() *Chitanda {
	ctd :=  &Chitanda{Engine: gin.New(), beanFactory: NewBeanFactory()}
	ctd.Use(ErrorHandler())
	config := InitConfig()
	ctd.beanFactory.setBean(config)
	return ctd
}

func (this *Chitanda) Start() {
	config := InitConfig()
	getCronTask().Start()
	this.Run(fmt.Sprintf(":%d", config.Server.Port))
}

func (this *Chitanda) Handle(httpMethod, relativePath string, handler interface{}) *Chitanda {
	if h:= Convert(handler);h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}

	return this
}

func (this *Chitanda) Responsible(f Responsible) *Chitanda {
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
		class.Build(this)
		this.beanFactory.inject(class)
		this.beanFactory.setBean(class)
	}
	return this
}


func (this *Chitanda) Task(expr string, f func()) *Chitanda {
	_, err := getCronTask().AddFunc(expr, f)
	if err != nil {
		log.Println(err)
	}
	return this
}