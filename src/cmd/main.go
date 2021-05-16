package main

import (
	"github.com/fyxemmmm/chitanda/chitanda"
	"github.com/fyxemmmm/chitanda/src/classes"
)

func main()  {
	chitanda.Inquisitive().
		Mount(
			classes.NewIndexClass(),
			classes.NewUserClass(),
			).
		Launch()
}

