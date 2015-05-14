package gch

import (
	"fmt"
)

// Anarchist - just forward all input

type Anarchist struct {
	// no data structures needed
}

func (a *Anarchist) Proc(in <-chan byte, out chan<- byte) {
	for {
		b := <-in      // grab new event
		fmt.Println(b) // print it
		out <- b       // forward it
	}
}
