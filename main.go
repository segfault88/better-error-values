package main

import (
	"fmt"
	"runtime/debug"
)

func a() {
	b()
}

func b() {
	c()
}

func c() {
	d()
}

func d() {
	fmt.Printf("\n\nhere is the stack:\n")
	debug.PrintStack()
}

func main() {
	a()

	c()
}
