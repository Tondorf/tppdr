package gio

import "fmt"
import "bytes"
import "strconv"
import "os/exec"

func GetActiveWindow() (windowID int) {

	cmd := exec.Command("xdotool", "getactivewindow")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error getting active window: ", err.Error())
	}

	i, _ := strconv.Atoi(out.String())
	return i

}
