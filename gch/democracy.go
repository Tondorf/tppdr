package gch

import (
	"fmt"
	"time"
)

const LEGISLATIVE_PERIOD = 2 // Ticks per second

type Democrat struct {
	values map[int]byte // serves as queue
}

func (d *Democrat) Proc(in <-chan byte, out chan<- byte) {
	d.values = make(map[int]byte)
	go d.delayedOut(out)
	for {
		b := <-in                   // grab new event
		d.values[len(d.values)] = b // append to the end of the queue
	}
}

func (d *Democrat) delayedOut(out chan<- byte) {
	for {
		time.Sleep(1000 / LEGISLATIVE_PERIOD * time.Millisecond)
		if len(d.values) > 0 { // only act if there are any votes
			v := democratize(d.values)
			fmt.Println("Democracy voted for", v)
			out <- v
			d.values = make(map[int]byte)
		} else {
			fmt.Println("no one voted")
		}
	}
}

func democratize(b map[int]byte) (maxKey byte) { // don't call with empty list!
	votes := make(map[byte]int)
	for _, v := range b {
		votes[v] += 1 // accumulate frequency of each Keycode
	}
	fmt.Println(votes)

	maxAmount := 0
	for key, amount := range votes { // determine most common Keycode
		if amount > maxAmount {
			maxAmount = amount
			maxKey = key
		}
	}

	return maxKey
}
