package main

import (
	"github.com/fyxemmmm/chitanda/chitanda"
	"github.com/fyxemmmm/chitanda/src/classes"
)

func main()  {
	chitanda.Inquisitive().
		Mount(
			"v1",
			classes.NewIndexClass(),
			classes.NewUserClass(),
			).
		Mount("v2",
			classes.NewUserClass()).
		Launch()
}


