package gch

import (
	"fmt"
	"time"
)

type Algo interface {
	Proc(in <-chan byte, out chan<- byte)
}

type Udi struct {
	values map[int]byte // serves as queue
}

func (u *Udi) delayedOut(out chan<- byte) {
	for {
		time.Sleep(1 * time.Second)
		for _, v := range u.values {
			fmt.Println(v)
			out <- v
		}
		u.values = make(map[int]byte)
	}
}

func (u *Udi) Proc(in <-chan byte, out chan<- byte) {
	u.values = make(map[int]byte)
	go u.delayedOut(out)
	for {
		b := <-in
		u.values[len(u.values)] = b
	}
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
