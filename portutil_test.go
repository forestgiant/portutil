package portutil

import (
	"net"
	"testing"
)

func TestVerify(t *testing.T) {
	port, err := Verify(8080)
	if err != nil {
		t.Errorf("Couldn't Verify port: %s. Error: %s", port, err)
	}
}

func TestVerifyHostPort(t *testing.T) {
	address, err := VerifyHostPort("127.0.0.1:8080")
	if err != nil {
		t.Errorf("Couldn't VerifyHostPort: %s. Error: %s", address, err)
	}
}

func TestGetUnique(t *testing.T) {
	_, err := GetUnique()
	if err != nil {
		t.Errorf("Err getting unique port", err)
	}
}

func TestPortTaken(t *testing.T) {
	uniqPort, _ := GetUnique()

	tests := []struct {
		port int
	}{
		{8090},
		{uniqPort},
	}

	for _, test := range tests {
		address := JoinHostPort("127.0.0.1", test.port)
		ln, err := net.Listen("tcp", address)
		if err != nil {
			t.Fatalf("Err creating listener", err)
		}
		defer ln.Close()

		address, err = VerifyHostPort(address)
		if err == nil {
			t.Errorf("Failed to detect port was taken. Error: %s", err)
		}
	}
}
