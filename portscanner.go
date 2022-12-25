package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPort scans the specified host and port and returns a boolean value indicating
// whether the port is open or closed, and an error if one occurred during the scan.
func ScanPort(host string, port int) (bool, error) {
	// Create a target address by combining the host and port
	target := fmt.Sprintf("%s:%d", host, port)

	// Create a TCP connection to the target address with a timeout of 5 seconds
	conn, err := net.DialTimeout("tcp", target, time.Second*5)
	if err != nil {
		// An error occurred, which means the port is probably closed
		return false, err
	}

	// No error occurred, which means the port is probably open
	defer conn.Close()
	return true, nil
}
