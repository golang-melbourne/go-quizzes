package main

import "fmt"

// START OMIT

type Gopher struct {
	Name string
}

func main() {
	orig := []Gopher{
		{"Whitaker"},
		{"Lindy"},
		{"Spike"},
	}

	second := make([]Gopher, len(orig))
	for i, a := range orig {
		second[i] = a
	}

	fmt.Printf("Values:   %v %v %v\n", orig[0], orig[1], orig[2])
	fmt.Printf("2nd Slice: %v %v %v\n", second[0], second[1], second[2])
}

// END OMIT
