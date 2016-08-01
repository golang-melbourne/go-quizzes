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
	gophers := []Gopher{
		{"Whitaker"},
		{"Lindy"},
		{"Spike"},
	}

	// START OMIT

	for _, g := range gophers {
		clone := g
		go clone.SayMyName()
	}

	// END OMIT

	time.Sleep(500 * time.Millisecond)

}
