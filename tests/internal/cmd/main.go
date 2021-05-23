package main

import (
	"fmt"
	"github.com/fyxemmmm/chitanda-gin/chitanda"
	"github.com/fyxemmmm/chitanda-gin/tests/internal/classes"
	"github.com/fyxemmmm/chitanda-gin/tests/internal/middlewares"
)

func main()  {
	mysqlHost := "localhost:3306"
	mysqlUsername := "root"
	mysqlUserPassword := "root"
	chitanda.Inquisitive().
		Joyful(chitanda.NewSqlXAdapter(mysqlHost, mysqlUsername, mysqlUserPassword)).
		Responsible(middlewares.NewUserMiddleware()).
		Earnest("v2",
			classes.NewUserClass()).
		Task("0/3 * * * * *", func() {
			fmt.Println("执行任务")
		}).
		Start()
}


