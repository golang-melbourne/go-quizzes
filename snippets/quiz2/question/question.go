package main

import (
	"fmt"
	"time"
)

// START OMIT

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

	for _, g := range gophers {
		go g.SayMyName()
	}

	time.Sleep(500 * time.Millisecond)

}

// END OMIT
