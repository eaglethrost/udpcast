package main

import (
	"fmt"
	"net"
)

/*
	Program to do udp broadcast with data of 1444 bytes
	First 4 byte offset is sequence of packet
*/

const (
	PKT_MAX_LEN = 1444
	PORT        = 5000
)

func main() {

	// 1. Create UDP Connection

	// addr = *net.UDPAddr
	// Get local address
	conn, err := net.Dial("udp", "192.168.2.20:5000")
	if err != nil {
		err = fmt.Errorf("can't dial")
	}
	addr := conn.LocalAddr().(*net.UDPAddr)
	conn.Close()

	udpl, err := net.ListenUDP("udp", addr)
	if err != nil {
		err = fmt.Errorf("listenudp :%s", addr.String())
	}
	defer udpl.Close()
	fmt.Printf("Connected...  %s\n", addr)

	// TESTING -----
	// 2. (TEST) Write to connection
	bcast_addr := &net.UDPAddr{
		IP:   net.IP{0xFF, 0xFF, 0xFF, 0xFF},
		Port: 5000,
	}
	msg := []byte("hello world")
	udpl.WriteToUDP(msg, bcast_addr)
	// TESTING ------

	// 2. Prepare UDP pakcet
	// 3. Broadcast packet
}
