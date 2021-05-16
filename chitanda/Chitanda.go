package chitanda

import "github.com/gin-gonic/gin"

type Chitanda struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Inquisitive() *Chitanda {
	return &Chitanda{Engine: gin.New()}
}

func (this *Chitanda) Launch() {
	this.Run(":8080")
}

func (this *Chitanda) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Chitanda {
	this.g.Handle(httpMethod, relativePath, handlers...)
	return this
}

func (this *Chitanda) Mount(group string, classes ...IClass) *Chitanda {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}
	return this
}

