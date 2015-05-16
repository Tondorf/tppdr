package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/Tondorf/tppdr/gch"
	"github.com/Tondorf/tppdr/gio"
	"github.com/Tondorf/tppdr/net"
)

func handleOutput(ch <-chan net.Key, windowID string) {
	for {
		v, ok := <-ch
		if ok {
			fmt.Println("ch: ", v)
		} else {
			time.Sleep(1 * time.Second)
		}

		// temporary workaround:
		// currently we are sending key events via syscall to xdotool
		// not nice, but hey: at least it works ;)
		gio.SendKey(windowID, v)
	}
}

func main() {
	fmt.Println("\033[1;37;44m -= tppdr Server =- \033[0m")
	var gamebin string
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter: Game to run")
		defaultGame := "/usr/games/bin/supertux"
		// equivalent to Python's `if os.path.exists(filename)`
		if _, err := os.Stat(defaultGame); err == nil {
			fmt.Println("... but found Default Game:\n    " + defaultGame)
			gamebin = defaultGame
		} else {
			fmt.Println("Default Game not found")
			os.Exit(1)
		}
	} else {
		gamebin = os.Args[1]
	}
	fmt.Println("awaiting connections ...")

	// start game
	cmd := exec.Command(gamebin)
	go cmd.Start()
	defer cmd.Wait()

	// determine window id for xdotool
	var windowID string = gio.GetActiveWindow()
	fmt.Println(windowID)

	// routing channels in order to use the ghc:
	chi := make(chan net.Key)
	cho := make(chan net.Key)

	// Governmental Algorithm for GCH
	// Should be interchangeable on-the-fly later
	go gch.Process(new(gch.Anarchist), chi, cho)
	//go gch.Process(new(gch.Democrat), chi, cho)

	// forward output to ... hum, well: to the output :p
	go handleOutput(cho, windowID)

	// run the server
	go net.Listen(1234, chi)

	fmt.Println("setup done")
}
