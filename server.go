package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/TondorfHacking/tppdr-gch"
	"github.com/TondorfHacking/tppdr-gio"
	"github.com/TondorfHacking/tppdr-net"
)

func handleOutput(ch <-chan byte) {
	for {
		v, ok := <-ch
		if ok {
			fmt.Println("ch: ", v)
		} else {
			time.Sleep(1 * time.Second)
		}

		// temporary workaround:
		// currently we are sending key events
		// via syscall to xdotool - not nice, but hey: at least it works ;)
		arg1 := os.Args[1]
		win, _ := strconv.ParseInt(arg1, 10, 64)
		gio.SendKey(int(win), v)
	}
}

func main() {
	fmt.Println("\033[1;37;44m -= tppdr Server =- \033[0m")
	fmt.Println("awaiting connections ...")

	// routing channels in order to use the ghc:
	chi := make(chan byte)
	cho := make(chan byte)

	// connect the gch using a given Algo:
	go gch.Process(new(gch.Udi), chi, cho)

	// forward output to ... hum, well: to the output :p
	go handleOutput(cho)

	// run the server
	err := net.Listen(1234, chi)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
	}

	fmt.Println("done")
}
