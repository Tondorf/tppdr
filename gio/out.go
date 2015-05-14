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
import "os/exec"

const xdotool = "/usr/bin/xdotool"

func SendKey(windowID string, keycode byte) (err error) {
	s := MapKey(keycode)
	if s != "" {
		//var cs *C.char = C.CString(bin + " key --window " + strconv.Itoa(window) + " --delay 25 " + s)
		cmd := exec.Command(xdotool, "key", "--window", windowID, "--delay 25 ", s)
		//_, err = C.cexec(cs)
		cmd.Start()
	}
	return
}

// wrap me if you can
func MapKey(keycode byte) string {
	return string(keycode)
}
