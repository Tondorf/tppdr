package gch // gch = game control heuristics

import (
	"github.com/Tondorf/tppdr/net"
)

type Algo interface {
	Proc(in <-chan net.Key, out chan<- net.Key)
}

func Process(a Algo, in <-chan net.Key, out chan<- net.Key) {
	a.Proc(in, out)
}

type Gamemode int

const (
	Anarchy Gamemode = iota
	Democracy
	Patriotism
)

var MODE Gamemode // = Anarchy

var in chan net.Key
var out chan net.Key

func SetMode(newM Gamemode) {
	MODE = newM
}
