package classes

import (
	"github.com/fyxemmmm/chitanda/chitanda"
	"github.com/gin-gonic/gin"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (this *IndexClass) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "index ok",
		})
	}
}

func (this *IndexClass) Build(chitanda *chitanda.Chitanda)  {
	// r.Handle
	chitanda.Handle("GET", "/", this.GetIndex())
}
