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
import "fmt"

import "github.com/Tondorf/tppdr/net"
import "github.com/rthornton128/goncurses"

const xdotool = "/usr/bin/xdotool"

func SendKey(windowID string, nkey net.Key) (err error) {
	key := mapKey(nkey)
	if key != "" {
		xcmd := xdotool + " key --delay 25 --window " + windowID + " " + key
		var cs *C.char = C.CString(xcmd)
		fmt.Println("xcmd:", xcmd)
		//cmd := exec.Command(xcmd)
		_, err = C.cexec(cs)
		//cmd.Start()
	}
	return
}

var Mapping = map[goncurses.Key]string{
	goncurses.KEY_UP:    "Up",
	goncurses.KEY_DOWN:  "Down",
	goncurses.KEY_LEFT:  "Left",
	goncurses.KEY_RIGHT: "Right",
}

// wrap me if you can
func mapKey(nkey net.Key) string {
	//fmt.Println(string(keycode))

	var cursesKey = nkey.K

	var keyStr = Mapping[cursesKey]

	return keyStr

	return goncurses.KeyString(nkey.K) // huehue
}
