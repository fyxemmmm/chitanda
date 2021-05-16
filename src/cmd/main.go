package main

import (
	"github.com/fyxemmmm/chitanda/chitanda"
	"github.com/fyxemmmm/chitanda/src/classes"
	"github.com/fyxemmmm/chitanda/src/middlewares"
)

func main()  {
	chitanda.Inquisitive().
		Responsible(middlewares.NewUserMiddleware()).
		Earnest(
			"v1",
			classes.NewIndexClass(),
			classes.NewUserClass(),
			).
		Earnest("v2",
			classes.NewUserClass()).
		Start()
}

