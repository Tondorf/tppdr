package main

// #include <stdlib.h>
//
// void cexec(char *c) {
//   int n = system(c);
//   free(c);
// }
import "C"
import "os"
import "fmt"

const xdotool = "/usr/bin/xdotool"

func main() {
	xcmd := xdotool + " key --delay 25 --window " + os.Args[1] + " " + os.Args[2]
	fmt.Println("xcmd:", xcmd)
	//cmd := exec.Command(xdotool, "key --delay 25 --window", os.Args[1], os.Args[2])
	//fmt.Println(cmd)
	//cmd.Run()
	var cs *C.char = C.CString(xcmd)
	//cmd := exec.Command(xcmd)
	out, err := C.cexec(cs)
	fmt.Println(out)
	fmt.Println(err)
}
