package portutil

import (
	"net"
	"testing"
)

func TestVerify(t *testing.T) {
	port, err := Verify(8080)
	if err != nil {
		t.Fatalf("Couldn't Verify port: %s. Error: %s", port, err)
	}
}

func TestVerifyHostPort(t *testing.T) {
	address, err := VerifyHostPort("127.0.0.1:8080")
	if err != nil {
		t.Fatalf("Couldn't VerifyHostPort: %s. Error: %s", address, err)
	}
}

func TestPortTaken(t *testing.T) {
	address := "127.0.0.1:8090"
	ln, err := net.Listen("tcp", address)
	if err != nil {
		t.Fatalf("Err creating listener", err)
	}
	defer ln.Close()

	address, err = VerifyHostPort(address)
	if err == nil {
		t.Fatalf("Failed to detect port was taken. Error: %s", err)
	}
}
