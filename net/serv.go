package net

import (
	"net"
	"strconv"
)

type Client struct {
	conn net.Conn
	// hier könnten auch channels zur rückkommunikation gebaut werden
}

type Net int

func (c Client) readfrom(ch chan<- byte) error {
	//defer close(ch)
	buf := make([]byte, 1024)
	for {
		n, err := c.conn.Read(buf)
		if err != nil || n == 0 {
			return err
		}
		for _, v := range buf[0:n] {
			ch <- v
		}
	}
return nil
}

func connHandler(con net.Conn, ch chan<- byte) {
	//defer con.Close()

	cl := Client{con}
	go cl.readfrom(ch)
}

func Listen(port int, ch chan<- byte) error {
	//addr, err := net.ResolveTCPAddr("ip4", ":"+string(port))
	//if err != nil {
	//	return err
	//}

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
