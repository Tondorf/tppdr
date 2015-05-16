package gio

// // some c helper again:
// // for the moment just call xdotool
// // replace with handler later on...
//
// //#include <stdlib.h>
//
// //void cexec(char *c) {
// //  system(c);
// //  free(c);
// //}
import "C"
import "fmt"
import "os/exec"
import "github.com/Tondorf/tppdr/net"
import "github.com/rthornton128/goncurses"

const xdotool = "/usr/bin/xdotool"

func SendKey(windowID string, nkey net.Key) (err error) {
	key := mapKey(nkey)
	if key != "" {
		//var cs *C.char = C.CString(xdotool + " key --window " + strconv.Itoa(window) + " --delay 25 " + s)
		xcmd := xdotool + " key --delay 25 --window " + windowID + " " + key
		fmt.Println(xcmd)
		cmd := exec.Command(xcmd)
		//_, err = C.cexec(cs)
		cmd.Start()
	}
	return
}

//Mapping := map[goncurses.Key]string{
//    KEY_LEFT: "Left",
//}

// wrap me if you can
func mapKey(nkey net.Key) string {
	//fmt.Println(string(keycode))

	//var cursesKey goncurses.Key
	//cursesKey = key.K

	//var keyStr string
	//keyStr = Mapping[cursesKey]

	//return keyStr

	return goncurses.KeyString(nkey.K) // huehue
}
