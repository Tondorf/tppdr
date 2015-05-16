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

const xdotool = "/usr/bin/xdotool"

func SendKey(windowID string, keycode byte) (err error) {
	key := MapKey(keycode)
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

// wrap me if you can
func MapKey(keycode byte) string {
	fmt.Println(string(keycode))
	return string(keycode)
}
