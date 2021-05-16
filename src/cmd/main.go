package main

import (
	"github.com/fyxemmmm/chitanda/src/chitanda"
	"github.com/fyxemmmm/chitanda/src/classes"
	"github.com/fyxemmmm/chitanda/src/middlewares"
)

func main()  {
	chitanda.Inquisitive().
		Joyful(chitanda.NewSqlXAdapter()).
		Responsible(middlewares.NewUserMiddleware()).
		//Earnest("v2",
		//	classes.NewUserClass()).
		Earnest("v1",
			classes.NewIndexClass()).
		Start()
}

