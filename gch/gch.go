package gch

type Algo interface {
	Proc(in <-chan byte, out chan<- byte)
}

func Process(a Algo, in <-chan byte, out chan<- byte) {
	a.Proc(in, out)
}

type Gamemode int

const (
	Anarchy Gamemode = iota
	Democracy
	Patriotism
)

var MODE Gamemode // = Anarchy

var in chan byte
var out chan byte

func SetMode(newM Gamemode) {
	MODE = newM
}
