package gio

// // some c helper again:
// // for the moment just call xdotool
// // replace with handler later on...
//
// #include <stdlib.h>
//
// void cexec(char *c) {
//   system(c);
//   free(c);
// }
import "C"
import "strconv"

const bin = "/usr/bin/xdotool"

func SendKey(window int, keycode byte) (err error) {
	s := MapKey(keycode)
	if s != "" {
		var cs *C.char = C.CString(bin + " key --window " + strconv.Itoa(window) + " --delay 25 " + s)
		_, err = C.cexec(cs)
	}
	return
}

// wrap me if you can
func MapKey(keycode byte) string {
	return string(keycode)
}
