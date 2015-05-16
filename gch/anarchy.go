package gch

import (
	"fmt"

	"github.com/Tondorf/tppdr/net"
)

// Anarchist - just forward all input

type Anarchist struct {
	// no data structures needed
}

func (a *Anarchist) Proc(in <-chan net.Key, out chan<- net.Key) {
	fmt.Println("ANARCHY!!!")
	for {
		b := <-in // grab new event
		//fmt.Println(b) // print it
		out <- b // forward it
	}
}
