package main

import (
	"fmt"
	"github.com/fyxemmmm/chitanda/src/chitanda"
	"github.com/fyxemmmm/chitanda/src/classes"
	"github.com/fyxemmmm/chitanda/src/middlewares"
)

func main()  {
	chitanda.Inquisitive().
		Joyful(chitanda.NewSqlXAdapter()).
		Responsible(middlewares.NewUserMiddleware()).
		Earnest("v2",
			classes.NewUserClass()).
		Task("0/3 * * * * *", func() {
			fmt.Println("执行任务")
		}).
		Start()
}



