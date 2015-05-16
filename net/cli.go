package net

import (
	"encoding/gob"
	"fmt"
	"net"
)

func SendTo(srv string, ch <-chan Key) error {
	conn, err := net.Dial("tcp", srv)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		k, ok := <-ch
		if ok == false {
			fmt.Println("channel error")
			return nil
		}
		enc := gob.NewEncoder(conn) // Encoder
		err := enc.Encode(k)
		if err != nil {
			fmt.Println("encode error:", err)
		}
	}
	return nil
}
