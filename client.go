package main

import (
	"fmt"
	"os"

	"github.com/Tondorf/tppdr/io"
	"github.com/Tondorf/tppdr/net"
	//	"log"
)

func main() {
	fmt.Println("\033[1;37;44m -= tppdr client =- \033[0m")
	if len(os.Args) < 2 {
		fmt.Println("missing expected parameter: 'host:port'")
		os.Exit(1)
	}
	serv := os.Args[1]
	fmt.Println("Connecting to", serv)

	ch := make(chan net.Key)
	go net.SendTo(serv, ch)

	//keyboard, cleanup := io.GetKeyboard()
	//defer cleanup()

	fmt.Println("just start kicking your keyboard 'round... (q to quit)\r")

	//keyboard.Listen(func(keycode byte) { ch <- keycode }, byte('q'))
	nckb := new(io.NCursesKeyboard)
	nckb.Listen(func(key net.Key) { ch <- key })

	fmt.Println("done")
}
