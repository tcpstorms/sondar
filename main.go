package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Specify the host and port to scan
	var host string
	fmt.Print("Enter the domain: ")
	fmt.Scanln(&host)

	var port int
	fmt.Print("Enter Port number: ")
	fmt.Scanln(&port)

	// Create a target address by combining the host and port
	target := fmt.Sprintf("%s:%d", host, port)

	// Create a TCP connection to the target address
	conn, err := net.DialTimeout("tcp", target, time.Second*5)
	if err != nil {
		// An error occurred, which means the port is probably closed
		fmt.Printf("Port %d is closed\n", port)
	} else {
		// No error occurred, which means the port is probably open
		fmt.Printf("Port %d is open\n", port)
		// Don't forget to close the connection when you're done
		conn.Close()
	}
}
