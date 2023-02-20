package gio

// // some c helper again:
// // for the moment just call xdotool
// // replace with handler later on...
//
// #include <stdlib.h>
//
// void cexec(char *c) {
//   int n = system(c); // get rid of warning, lol
//   free(c);
// }
import "C"
import (
	"fmt"
	"github.com/Tondorf/tppdr/common"
)

const xdotool = "/usr/bin/xdotool"

func SendKey(windowID string, ge common.GameEvent) (err error) {
	key := mapKey(ge.Key)
	if key != "" {
		var xcmd string
		if ge.Typ == common.Press {
			xcmd = xdotool + " keydown --window " + windowID + " " + key
		} else if ge.Typ == common.Release {
			xcmd = xdotool + " keyup --window " + windowID + " " + key
		}
		var cs *C.char = C.CString(xcmd)
		fmt.Println("xcmd:", xcmd)
		//cmd := exec.Command(xcmd)
		_, err = C.cexec(cs)
		//cmd.Start()
	}
	return
}

var Mapping = map[string]string{
	"ArrowUp":    "Up",
	"ArrowDown":  "Down",
	"ArrowLeft":  "Left",
	"ArrowRight": "Right",
	"Enter":      "Return",
	"Escape":     "Escape",
	" ":          "space",
}

// wrap me if you can
func mapKey(input string) (result string) {
	var keyStr = Mapping[input]
	return keyStr
}
