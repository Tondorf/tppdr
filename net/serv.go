package net

import (
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"strconv"
)

type Client struct {
	conn net.Conn
	// hier könnten auch channels zur rückkommunikation gebaut werden
}

type Net int

func (c Client) readfrom(ch chan<- Key) error {
	//defer close(ch)
	//buf := make([]byte, 1024)
	dec := gob.NewDecoder(c.conn) // Decoder
	for {
		//n, err := c.conn.Read(buf)
		//if err != nil || n <= 0 {
		//	return err
		//}
		//for _, v := range buf[0:n] {
		//	ch <- v
		//}
		var k Key
		err := dec.Decode(&k)
		if err == io.EOF {
			fmt.Println("decode EOF:", err)
			return nil
		} else if err != nil {
			fmt.Println("decode error:", err)
			return err
		}
		ch <- k // send key to the channel
	}
	return nil
}

func connHandler(con net.Conn, ch chan<- Key) {
	//defer con.Close()

	cl := Client{con}
	go cl.readfrom(ch)
}

func Listen(port int, ch chan<- Key) error {

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go connHandler(conn, ch)
	}
	return nil
}
