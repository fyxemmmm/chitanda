# chitanda
Chitanda is a web scaffold which is based on gin.

## ⌛ Installation
`go get https://github.com/fyxemmmm/chitanda-gin`

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



// 类
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
// 每个类需要实现Build方法
func (this *UserClass) Build(chitanda *chitanda.Chitanda)  {
	chitanda.Handle("GET", "/user-list", this.UserList)
}

// start server
func main()  {
    chitanda.Inquisitive().
    Earnest("v2", NewUserClass()).
    Start()
}
```

#### single model returns
```go
// 返回  chitanda.Model, model需要实现 String方法
func (this *UserClass) UserDetail(ctx *gin.Context) chitanda.Model {
	user := &models.UserModel{}
	return user
}
```

#### multi models return
```go
// 返回切片, 最后调用chitanda.ToModels(users)进行转换即可
func (this *UserClass) UserList(ctx *gin.Context) chitanda.Models {
    users := []*models.UserModel{
        {UserId: 101, UserName: "feixiang101"},
        {UserId: 102, UserName: "feixiang102"},
    }
    return chitanda.ToModels(users)
}
```

#### middlewares
```go
chitanda.Inquisitive().
    Responsible(middlewares.NewUserMiddleware()). // 中间件
    Earnest("v2",
        classes.NewUserClass()).
    }).
    Start()

//中间件需实现`onRequest`方法
type UserMiddleware struct {}

func NewUserMiddleware() *UserMiddleware {
return &UserMiddleware{}
}

func (this *UserMiddleware) OnRequest(ctx *gin.Context) error {
    // your logic
    return nil
}


```

### ⚡ async task
```go
func (this *UserClass) UserDetail(ctx *gin.Context) chitanda.Model {
	user := &models.UserModel{}

	// 这里是异步任务, 第二个参数是callback, 执行完成AddFavour会被hook
	chitanda.Task(this.AddFavour, func() {
		fmt.Println("doing callback")
	}, "params0")
	return user
}

func (this *UserClass) AddFavour(params ...interface{}) {
	fmt.Println(params[0])
	fmt.Println("add favor")
	time.Sleep(time.Second * 3)
}
```
#### crontab task
```go
chitanda.Inquisitive().
    Responsible(middlewares.NewUserMiddleware()).
    Earnest("v2",
        classes.NewUserClass()).
    Task("0/3 * * * * *", func() {  // 与linux的定时任务写法一致
        fmt.Println("执行定时任务")
    }).
    Start()
```

### 🍭 dependency injection & sqlx orm
```go
// 启动服务的时候, Joyful方法允许将对象进行注入
// 注入的对象需要是指针对象且初始化的时候没有赋值, 如果有值, 不会进行修改
// 这里以内置的chitanda.sqlx为例
type UserClass struct {
    *chitanda.SqlXAdapter 
}

mysqlHost := "localhost:3306"
mysqlUsername := "root"
mysqlUserPassword := "root"
chitanda.Inquisitive().
	// 这里注入sqlx对象
    Joyful(chitanda.NewSqlXAdapter(mysqlHost, mysqlUsername, mysqlUserPassword)).
    Earnest("v2",
        classes.NewUserClass()).
    Start()


func (this *UserClass) UserDetail(ctx *gin.Context) chitanda.Model {
    user := &models.UserModel{}
    sql := "select id, name, age, email from my.user where id = ?"
    err = this.GetContext(ctx, user, sql, user.UserId) // 直接可以使用
    if err != nil {
        chitanda.Error(err)
    }
    return user
}
```

>  ⚠️ If you want to inject your own orm adapter, you should inject it by yourself.

---
#### configuration file
- put the configuration file named "application.yaml" In your main directory
- user config will inject into your controller
- while compiling && go version < 1.16, it would not put the static file into executable file, this should be concerned.
```yaml
# 这是系统配置, 启动服务的时候自定义的端口号, 不填默认8080
# application.yaml
server:
  port: 8081

# 用户自定义配置, 必须在config这个层级下
config:
  user:
    age: 100
```

user config file useage
```go
type UserClass struct {
	// 这个会去取配置文件 user.age的内容, 类型是*chitanda.Value
	Age *chitanda.Value `prefix:"user.age"`  
}

func (this *UserClass) UserTest(ctx *gin.Context) string {
    return "my age is: " + this.Age.String()  // oh my age is 100!
}
```

>  ⚠️ you should compile the application.yaml into the executable file.   
&emsp;&nbsp; here is a lib: https://github.com/mjibson/esc
1. esc -o static.go application.yaml
2. go build main.go static.go

✅ now your configuration file would be effective.

---
#### ❤ you can take the complete example in tests/internal
