package main

import (
	"fmt"
	"net"
	"time"
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
	udpl.Close()

	fmt.Printf("Connected...  %s\n", addr)

	// TESTING -----
	// 2. (TEST) Write to connection
	bcast_addr := &net.UDPAddr{
		IP:   net.IP{0xFF, 0xFF, 0xFF, 0xFF},
		Port: 5000,
	}
	msg := []byte("siuuu")
	for {
		time.Sleep(2 * time.Second)
		udpl.WriteToUDP(msg, bcast_addr)
	}
	// TESTING ------

	// 2. Prepare UDP pakcet

	// 3. Broadcast packet
}

// package main

// import (
// 	"net"
// )

// func main() {
// 	pc, err := net.ListenPacket("udp4", ":5000")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer pc.Close()

// 	addr, err := net.ResolveUDPAddr("udp4", "192.168.7.255:8829")
// 	if err != nil {
// 		panic(err)
// 	}

// 	pc.WriteTo([]byte("data to transmit"), addr)
// }
