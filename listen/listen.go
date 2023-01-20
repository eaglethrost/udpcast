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
	fmt.Println("listen script")

	// 1. TAP INTO CONNECTION using server port information
	listenAddr := &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: PORT}
	udpconn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		panic(err)
	}
	defer udpconn.Close()

	// 2. LISTEN TO BROADCASTS
	buf := make([]byte, 1444)
	n, addr, err := udpconn.ReadFrom(buf)
	if err != nil {
		panic(err)
	}

	// 3. PRINT BROADCASTED MESSAGE
	fmt.Printf("%s sent this: %s\n", addr, string(buf[:n]))
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
