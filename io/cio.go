package io

/*
 * One mayor drawback:
 *  The present solution handles key presses only!
 *  We need keydown/keyup events!!!
 */

// // Let's begin with some pure c to get the basic keyboard input
//
// #include "stdio.h"
// #include "stdlib.h"
//
// int cgetchar() {
//   return getchar();
// }
//
// // Set the terminal into raw mode, deactivate echos
// void csetraw() {
//   int _ = system("/bin/stty raw -echo");
// }
//
// // Set the terminal back into cooked mode, enable echos
// void csetcooked() {
//   int _ = system("/bin/stty cooked echo");
// }
//
// // go for it:
import "C"

func setRaw() {
	C.csetraw()
	//exec.Command("/bin/stty raw -echo").Run()  // should work, but it doesnt
}

func setCooked() {
	C.csetcooked()
	//exec.Command("/bin/stty", "cooked", "echo").Run()  // dito
}

func getChar() (int, error) {
	c, e := C.cgetchar()
	if e != nil {
		return -1, e
	}
	return int(c), nil
}
