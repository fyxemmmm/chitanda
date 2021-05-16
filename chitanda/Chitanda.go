package chitanda

import "github.com/gin-gonic/gin"

type Chitanda struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Inquisitive() *Chitanda {
	return &Chitanda{Engine: gin.New()}
}

func (this *Chitanda) Start() {
	this.Run(":8080")
}

func (this *Chitanda) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Chitanda {
	this.g.Handle(httpMethod, relativePath, handlers...)
	return this
}

func (this *Chitanda) Responsible(f Responsible) *Chitanda{
	this.Use(func(context *gin.Context) {
		err := f.OnRequest()
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		}else {
			context.Next()
		}
	})
	return this
}



func (this *Chitanda) Earnest(group string, classes ...IClass) *Chitanda {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}
	return this
}

