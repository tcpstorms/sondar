package main

import (
	"fmt"

	"portscanner" // Import the scanner package
)

func main() {
	// Read the host and port from user input
	var host string
	fmt.Print("Enter the domain: ")
	_, err := fmt.Scanln(&host)
	if err != nil {
		fmt.Println("Error reading host:", err)
		return
	}

	var port int
	fmt.Print("Enter Port number: ")
	_, err = fmt.Scanln(&port)
	if err != nil {
		fmt.Println("Error reading port:", err)
		return
	}

	// Call the ScanPort function from the scanner package to scan the specified port
	status, err := scanner.ScanPort(host, port)
	if err != nil {
		fmt.Println("Error scanning port:", err)
		return
	}

	// Print the result of the port scan
	if status {
		fmt.Printf("Port %d is open\n", port)
	} else {
		fmt.Printf("Port %d is closed\n", port)
	}
}
