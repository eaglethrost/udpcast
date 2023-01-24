package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
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

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

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
	// msg := []byte("hello world")
	// udpl.WriteToUDP(msg, bcast_addr)
	// TESTING ------

	// 2. Prepare UDP pakcet
	bcast_addr := &net.UDPAddr{
		IP:   net.IP{0xFF, 0xFF, 0xFF, 0xFF},
		Port: 5000,
	}
	// 2.1 Get data
	fileBytes, err := os.ReadFile("test.zip")
	if err != nil {
		fmt.Printf("Error opening file \n")
	}

	// 2.2 Divide data into packets of 1444 bytes
	fileLen := len(fileBytes)
	no_of_pkts := fileLen / 1440
	fmt.Printf("Size: %d bytes Packets: %d\n", fileLen, no_of_pkts)

	for i := 0; i <= no_of_pkts; i++ {
		// 2.3 Set first 4 bytes as sequence number
		pkt := make([]byte, 1444)
		binary.BigEndian.PutUint16(pkt[:2], uint16(i+1))
		binary.BigEndian.PutUint16(pkt[2:4], uint16(no_of_pkts))

		// 2.4 Merge bytes
		edge := Min(1440*i+1440, fileLen)
		start := 1440 * i
		fmt.Printf("Seq %d) %d : %d\n", i, start, edge)
		copy(pkt[4:], fileBytes[start:edge])

		// 3. Broadcast packet
		udpl.WriteToUDP(pkt, bcast_addr)

		time.Sleep(time.Millisecond * 200)
		fmt.Printf("Sent packet:")
		// fmt.Println(pkt)
	}
	fmt.Println("Done sending")
}
