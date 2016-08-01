package main

import (
	"fmt"
	"time"
)

type Gopher struct {
	Name string
}

func (g *Gopher) SayMyName() {
	fmt.Println("My name is", g.Name)
}

func main() {

	// START OMIT

	ptrs := []*Gopher{
		&Gopher{"Whitaker"},
		&Gopher{"Lindy"},
		&Gopher{"Spike"},
	}

	for _, g := range ptrs {
		go g.SayMyName()
	}

	time.Sleep(500 * time.Millisecond)

	// END OMIT
}
