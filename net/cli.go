package net

import (
	"encoding/gob"
	"fmt"
	"net"
)

func SendTo(srv string, ch <-chan Key) error {
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		fmt.Println("Connection error:", err)
		return err
	}
	defer conn.Close()
	enc := gob.NewEncoder(conn) // Encoder

	for {
		k, ok := <-ch
		if ok == false {
			fmt.Println("channel error")
			return nil
		}
		err := enc.Encode(k)
		if err != nil {
			fmt.Println("encode error:", err)
		}
	}
	return nil
}
