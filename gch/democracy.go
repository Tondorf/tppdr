package gch

import (
	"fmt"
	"time"

	"github.com/Tondorf/tppdr/net"
)

// Democrat - accumulate Key presses and send the most desired one to the game

const LEGISLATIVE_PERIOD = 2 // Ticks per second

type Democrat struct {
	values map[int]net.Key // serves as queue
}

func (d *Democrat) Proc(in <-chan net.Key, out chan<- net.Key) {
	d.values = make(map[int]net.Key)
	go d.delayedOut(out)
	for {
		b := <-in                   // grab new event
		d.values[len(d.values)] = b // append to the end of the queue
	}
}

func (d *Democrat) delayedOut(out chan<- net.Key) {
	for {
		time.Sleep(1000 / LEGISLATIVE_PERIOD * time.Millisecond)
		if len(d.values) > 0 { // only act if there are any votes
			v := democratize(d.values)
			fmt.Println("Democracy voted for", v)
			out <- v
			d.values = make(map[int]net.Key)
		} else {
			fmt.Println("no one voted")
		}
	}
}

func democratize(b map[int]net.Key) (maxKey net.Key) { // don't call with empty list!
	votes := make(map[net.Key]int)
	for _, v := range b {
		votes[v] += 1 // accumulate frequency of each Keycode
	}
	fmt.Println(votes)

	var maxAmount int
	for key, amount := range votes { // determine most common Keycode
		if amount > maxAmount {
			maxAmount = amount
			maxKey = key
		}
	}

	return
}
