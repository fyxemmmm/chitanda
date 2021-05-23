# chitanda
<h1>
<img width="220" height="230"  align=right src="https://suki.fyxemmmm.cn/wp-content/themes/my/chitanda.png" />
</h1>

Chitanda is a web scaffold which is based on gin.


## ⌛ Installation
`https://github.com/fyxemmmm/chitanda-gin`

### 🔥 quick start
```go
package main

import (
	"github.com/fyxemmmm/chitanda-gin/chitanda"
	"github.com/gin-gonic/gin"
)

// 模型
type UserModel struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

func (this *UserModel) String() string {
	return "usermodel"
}



// 类似java中的类
type UserClass struct {}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserList(ctx *gin.Context) chitanda.Models {
	users := []*UserModel{
		{UserId: 888, UserName: "feixiang1209"},
		{UserId: 666, UserName: "fyxemmmmmmmm"},
	}
	return chitanda.ToModels(users)
}

// handle http请求
func (this *UserClass) Build(chitanda *chitanda.Chitanda)  {
	chitanda.Handle("GET", "/user-list", this.UserList)
}

```

