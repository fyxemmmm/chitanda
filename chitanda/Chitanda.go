package chitanda

import "github.com/gin-gonic/gin"

type Chitanda struct {
	*gin.Engine
}

func Inquisitive() *Chitanda {
	return &Chitanda{Engine: gin.New()}
}

func (this *Chitanda) Launch() {
	this.Run(":8080")
}

func (this *Chitanda) Mount(classes ...IClass) *Chitanda {
	for _, class := range classes {
		class.Build(this)
	}
	return this
}