package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/Tondorf/tppdrlib/gch"
	"github.com/Tondorf/tppdrlib/gio"
	"github.com/Tondorf/tppdrlib/net"
)

func handleOutput(ch <-chan byte, wid int) {
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
		//arg1 := os.Args[1]
		//win, _ := strconv.ParseInt(arg1, 10, 64)
		gio.SendKey(wid, v)
	}
}

func main() {
	fmt.Println("\033[1;37;44m -= tppdr Server =- \033[0m")
	if len(os.Args) < 2 {
		fmt.Println("missing paramter: game to run")
		os.Exit(1)
	}
	fmt.Println("awaiting connections ...")

	// start game
	gamebin := os.Args[1]
	cmd := exec.Command(gamebin)
	go cmd.Start()
	defer cmd.Wait()

	var wid int = gio.GetActiveWindow()

	// routing channels in order to use the ghc:
	chi := make(chan byte)
	cho := make(chan byte)

	// Governmental Algorithm for GCH
	// Should be interchangeable on-the-fly later
	go gch.Process(new(gch.Democrat), chi, cho)

	// forward output to ... hum, well: to the output :p
	go handleOutput(cho, wid)

	// run the server
	go net.Listen(1234, chi)
	//err := net.Listen(1234, chi)
	//if err != nil {
	//	fmt.Println("Error listening: ", err.Error())
	//}

	fmt.Println("setup done")
}
