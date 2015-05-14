package gch

import (
	"fmt"
)

type Anarchist struct {
}

func (a *Anarchist) Proc(in <-chan byte, out chan<- byte) {
	for {
		b := <-in      // grab new event
		fmt.Println(b) // print it
		out <- b       // forward it
	}
}
