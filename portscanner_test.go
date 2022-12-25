package portscanner_test

import (
	"net"
	"testing"
	"time"
  
  "portscanner" // Import the scanner package
)

func TestScanPort(t *testing.T) {
	// Test a closed port
	if err := scanPort("localhost", 8080); err == nil {
		t.Error("Expected error for closed port, got nil")
	}

	// Test an open port
	if err := scanPort("localhost", 80); err != nil {
		t.Error("Expected nil error for open port, got", err)
	}
}

func TestScanPortTimeout(t *testing.T) {
	// Test a port that takes longer than the timeout to respond
	if err := scanPort("localhost", 8080, time.Millisecond*100); err == nil {
		t.Error("Expected error for port with timeout, got nil")
	}
}

func TestScanPortInvalidHost(t *testing.T) {
	// Test an invalid hostname
	if err := scanPort("invalid.hostname", 80); err == nil {
		t.Error("Expected error for invalid hostname, got nil")
	}
}

func TestScanPortInvalidPort(t *testing.T) {
	// Test an invalid port number
	if err := scanPort("localhost", -1); err == nil {
		t.Error("Expected error for invalid port, got nil")
	}
}
