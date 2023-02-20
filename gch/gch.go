package gch // Package gch = game control heuristics
import "github.com/Tondorf/tppdr/common"

type Algo interface {
	Proc(in <-chan common.BrowserEvent, out chan<- common.GameEvent)
}

func Process(a Algo, in <-chan common.BrowserEvent, out chan<- common.GameEvent) {
	a.Proc(in, out)
}

type Gamemode int

const (
	Anarchy Gamemode = iota
	Democracy
	Patriotism
)

var MODE Gamemode // = Anarchy

var in chan common.BrowserEvent
var out chan common.GameEvent

func SetMode(newM Gamemode) {
	MODE = newM
}
