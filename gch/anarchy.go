package gch

import (
	"fmt"
	"github.com/Tondorf/tppdr/common"
)

type inputs map[string]bool

// Anarchist - just forward all input

type Anarchist struct {
	controls map[string]inputs // IP -> inputs
}

func (a *Anarchist) Proc(in <-chan common.BrowserEvent, out chan<- common.GameEvent) {
	fmt.Println("ANARCHY!!!")
	a.controls = make(map[string]inputs)
	for {
		be := <-in // grab new event

		// new presses/releases (in this tick)
		var newPresses = make([]string, 0)
		var newReleases = make([]string, 0)

		if _, prs := a.controls[be.Origin]; !prs {
			a.controls[be.Origin] = make(inputs, 0)
		}
		if be.Typ == common.Press {
			// iff state changed from false to true: trigger press event
			if a.controls[be.Origin][be.Key] == false {
				newPresses = append(newPresses, be.Key)
			}
			a.controls[be.Origin][be.Key] = true
		} else if be.Typ == common.Release {
			// iff state changed from true to false: trigger release event
			if a.controls[be.Origin][be.Key] == true {
				newReleases = append(newReleases, be.Key)
			}
			a.controls[be.Origin][be.Key] = false
		}

		//var actual_cmds = make([]string, 0)
		//for _, inps := range a.controls {
		//	for input, pressed := range inps {
		//		if pressed {
		//			actual_cmds = append(actual_cmds, input)
		//		}
		//	}
		//}

		for _, rel := range newReleases {
			ge := common.GameEvent{
				Key: rel,
				Typ: common.Release,
			}
			out <- ge
		}

		for _, press := range newPresses {
			ge := common.GameEvent{
				Key: press,
				Typ: common.Press,
			}
			out <- ge
		}

	}
}
