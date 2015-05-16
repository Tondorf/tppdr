package io

import "fmt"
import "github.com/Tondorf/tppdr/net"
import "github.com/rthornton128/goncurses"

type NCursesKeyboard int

func (n *NCursesKeyboard) Listen(handler func(net.Key)) {
	stdscr, err := goncurses.Init()
	if err != nil {
		fmt.Println(err)
	}
	//stdscr.Clear()
	stdscr.Keypad(true)
	for {
		key := stdscr.GetChar()
		if key == 'q' {
			break
		}
		handler(net.Key{K: key})
	}
}
