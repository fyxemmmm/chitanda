package chitanda

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				context.AbortWithStatusJSON(400, gin.H{"error": e})
			}
		}()
		context.Next()
	}
}

func Error(err error, msgs ...string)  {
	if err == nil {
		return
	}else {
		errMsg := err.Error()
		if len(msgs) > 0 {
			errMsg = msgs[0]
		}
		panic(errMsg)
	}
}
