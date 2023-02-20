package main

import (
	"fmt"
	"github.com/Tondorf/tppdr/common"
	"os"
	"os/exec"
	"time"

	"github.com/Tondorf/tppdr/gch"
	"github.com/Tondorf/tppdr/gio"
	"github.com/Tondorf/tppdr/web"
)

func handleOutput(ch <-chan common.GameEvent, windowID string) {
	for {
		v, ok := <-ch
		if ok {
			//fmt.Println("ch: ", v)
			// temporary workaround:
			// currently we are sending key events via syscall to xdotool
			// not nice, but hey: at least it works ;)
			gio.SendKey(windowID, v)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	fmt.Println("\033[1;37;44m -= tppdr Server =- \033[0m")
	var gamebin string
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter: Game to run")
		defaultGame := "/usr/games/supertux2"
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
	chi := make(chan common.BrowserEvent)
	cho := make(chan common.GameEvent)

	// start webserver
	webserver := web.NewWebserver(chi)
	go webserver.Listen()

	// Governmental Algorithm for GCH
	// Should be interchangeable on-the-fly later
	go gch.Process(new(gch.Anarchist), chi, cho)
	//go gch.Process(new(gch.Democrat), chi, cho)

	// forward output to ... hum, well: to the output :p
	go handleOutput(cho, windowID)

	fmt.Println("setup done")
}
