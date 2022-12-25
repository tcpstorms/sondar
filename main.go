package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Declare variables to hold the host and port to scan
	var host string
	var port int

	// Read the host and port from user input
	fmt.Print("Enter the domain: ")
	_, err := fmt.Scanln(&host)
	if err != nil {
		fmt.Println("Error reading host:", err)
		return
	}

	fmt.Print("Enter Port number: ")
	_, err = fmt.Scanln(&port)
	if err != nil {
		fmt.Println("Error reading port:", err)
		return
	}

	// Create a target address by combining the host and port
	target := fmt.Sprintf("%s:%d", host, port)

	// Create a TCP connection to the target address with a timeout of 5 seconds
	conn, err := net.DialTimeout("tcp", target, time.Second*5)
	if err != nil {
		// An error occurred, which means the port is probably closed
		fmt.Printf("Port %d is closed\n", port)
	} else {
		// No error occurred, which means the port is probably open
		fmt.Printf("Port %d is open\n", port)
		// Don't forget to close the connection when you're done
		defer conn.Close()
	}
}
