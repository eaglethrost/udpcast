package main

import (
	"fmt"
	"net"
)

const (
	PKT_MAX_LEN = 1444
	PORT        = 5000
)

func main() {
	// 1. TAP INTO CONNECTION using server port information
	listenAddr := &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: PORT}
	udpconn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		panic(err)
	}
	defer udpconn.Close()
	fmt.Printf("listen script on: \n")

	// 2. LISTEN TO BROADCASTS
	var totalFile []byte
	for {
		buf := make([]byte, 1444)
		n, addr, err := udpconn.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
		// 3. PROCESS BROADCASTED MESSAGE
		fmt.Printf("%s sent this: %d\n", addr, n)
		fmt.Println(buf)
		totalFile = append(totalFile, buf...)
	}
}

// package main

// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	pc, err := net.ListenPacket("udp4", ":5000")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer pc.Close()

// 	buf := make([]byte, 1024)
// 	n, addr, err := pc.ReadFrom(buf)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("%s sent this: %s\n", addr, string(buf[:n]))
// }
