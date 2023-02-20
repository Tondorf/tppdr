package gch

import (
	"fmt"
	"github.com/Tondorf/tppdr/common"
)

// Anarchist - just forward all input

type Anarchist struct {
	// no data structures needed
}

func (a *Anarchist) Proc(in <-chan common.BrowserEvent, out chan<- common.GameEvent) {
	fmt.Println("ANARCHY!!!")
	for {
		be := <-in // grab new event
		ge := common.GameEvent{
			Key: be.Key,
			Typ: be.Typ,
		}
		out <- ge // forward it
	}
}
