package main

import "fmt"

type Gopher struct {
	Name string
}

func main() {
	orig := []Gopher{
		{"Whitaker"},
		{"Lindy"},
		{"Spike"},
	}

	// START OMIT
	ptrs := make([]*Gopher, len(orig))
	for i, _ := range orig {
		ptrs[i] = &orig[i]
	}
	// END OMIT

	fmt.Printf("Values:   %v %v %v\n", orig[0], orig[1], orig[2])
	fmt.Printf("Pointers: %v %v %v\n", ptrs[0], ptrs[1], ptrs[2])
}
