package gio

import "fmt"
import "time"
import "bytes"
import "strings"

import "os/exec"

func GetActiveWindow() (windowID string) {

	time.Sleep(1000 * time.Millisecond) // wait for game window to open

	cmd := exec.Command("xdotool", "getactivewindow")
	//cmd := exec.Command("xdotool", "getwindowfocus")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	// TrimSpace removes newlines at the end
	return strings.TrimSpace(out.String())
}

func SelectWindow() (windowID string) {

	fmt.Println("Plase click in the Game window.")
	cmd := exec.Command("xdotool", "selectwindow")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	// TrimSpace removes newlines at the end
	return strings.TrimSpace(out.String())
}
