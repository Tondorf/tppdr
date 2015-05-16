package net

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/rthornton128/goncurses"
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
			goncurses.End()
			return nil
		}
		err := enc.Encode(k)
		if err != nil {
			fmt.Println("encode error:", err)
			goncurses.End()
			return nil
		}
	}
	return nil
}
