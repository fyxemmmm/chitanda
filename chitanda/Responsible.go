package chitanda

import "github.com/gin-gonic/gin"

type Responsible interface {
	OnRequest(ctx *gin.Context) error
}