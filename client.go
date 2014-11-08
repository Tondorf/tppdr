package main

import (
	"fmt"
	"os"

	"github.com/TondorfHacking/tppdr-io"
	"github.com/TondorfHacking/tppdr-net"
	//	"log"
)

func main() {
	fmt.Println("tppdr client")
	fmt.Println("============")
	s := os.Args[1]
	fmt.Println("Connecting to ", s)

	ch := make(chan byte)
	go net.SendTo(os.Args[1], ch)

	k, d := io.GetKeyboard()
	defer d()

	fmt.Println("just start kicking your keyboard 'round... (q to quit)\r")

	k.Listen(func(keycode byte) { ch <- keycode }, byte('q'))

	fmt.Println("done")
}
