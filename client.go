package main

import (
	"fmt"
	"os"

	"github.com/Tondorf/tppdr/io"
	"github.com/Tondorf/tppdr/net"
	//	"log"
)

func main() {
	fmt.Println("tppdr client")
	fmt.Println("============")
	if len(os.Args) < 2 {
		fmt.Println("missing expected paramter 'host:port'")
		os.Exit(1)
	}
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
