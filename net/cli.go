package net

import "net"
import "fmt"

func SendTo(srv string, ch <-chan byte) error {
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		return err
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		v, ok := <-ch
		if ok == false {
			fmt.Println("channel error")
			return nil
		}
		buf[0] = v
		n, err := conn.Write(buf[0:1])
		if err != nil || n == 0 {
			fmt.Println(err)
			return err
		}
	}
return nil
}
