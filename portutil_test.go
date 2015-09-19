package portutil

import (
	"net"
	"testing"
)

func TestVerifyTCP(t *testing.T) {
	testPort := 8000

	port, err := VerifyTCP(testPort)
	if err != nil {
		t.Errorf("Couldn't Verify port: %d. Error: %s", port, err)
		return
	}

	if port == 0 {
		t.Errorf("Couldn't Verify port: %d. Error: %s", port, err)
		return
	}

	if port != testPort {
		t.Errorf("testPort: %d should equal port: %d", testPort, port)
		return
	}
}

func TestVerifyUDP(t *testing.T) {
	testPort := 8000

	port, err := VerifyUDP(testPort)
	if err != nil {
		t.Errorf("Couldn't Verify port: %d. Error: %s", port, err)
		return
	}

	if port == 0 {
		t.Errorf("Couldn't Verify port: %d. Error: %s", port, err)
		return
	}

	if port != testPort {
		t.Errorf("testPort: %d should equal port: %d", testPort, port)
		return
	}
}

func TestVerify(t *testing.T) {
	testPort := 9000

	tests := []struct {
		netProto string
	}{
		{udp},
		{tcp},
	}
	for _, test := range tests {
		port, err := Verify(test.netProto, testPort)
		if err != nil {
			t.Errorf("Couldn't Verify port on UDP: %s. Error: %s", port, err)
			return
		}

		if port == 0 {
			t.Errorf("Couldn't Verify port on UDP: %d. Error: %s", port, err)
			return
		}

		if port != testPort {
			t.Errorf("testPort: %d should equal UDP port: %d", testPort, port)
			return
		}
	}
}

func TestVerifyHostPort(t *testing.T) {
	testAddr := "127.0.0.1:9080"
	tests := []struct {
		netProto string
	}{
		{udp},
		{tcp},
	}
	for _, test := range tests {
		addr, err := VerifyHostPort(test.netProto, testAddr)
		if err != nil {
			t.Errorf("Couldn't VerifyHostPort: %s. Error: %s", addr, err)
			return
		}

		if addr == "" {
			t.Errorf("Couldn't VerifyHostPort: %s. Error: %s", addr, err)
			return
		}

		if addr != testAddr {
			t.Errorf("testPort: %d should equal port: %d", testAddr, addr)
			return
		}
	}
}

func TestGetUniqueTCP(t *testing.T) {
	port, err := GetUniqueTCP()
	if err != nil || port == 0 {
		t.Fatalf("Err getting unique port TCP: %s", err)
	}
}

func TestGetUniqueUDP(t *testing.T) {
	port, err := GetUniqueUDP()
	if err != nil || port == 0 {
		t.Fatalf("Err getting unique port UDP: %s", err)
	}
}

func TestGetUnique(t *testing.T) {
	tests := []struct {
		netProto string
	}{
		{udp},
		{tcp},
	}
	for _, test := range tests {
		port, err := GetUnique(test.netProto)
		if err != nil || port == 0 {
			t.Fatalf("%s: Err getting unique port: %s", test.netProto, err)
		}
	}
}

// This test creates a listener at a port and then calls Verify()
func TestPortTaken(t *testing.T) {
	tests := []struct {
		netProto string
	}{
		{udp},
		{tcp},
	}

	for _, test := range tests {
		port, _ := GetUnique(test.netProto)

		// Create a listener
		var err error
		var ln interface{}

		switch test.netProto {
		case udp:
			ln, err = newListenerUDP(port)

			if err != nil {
				t.Errorf("Err creating UDP listener", err)
			}
			defer ln.(*net.UDPConn).Close()
		case tcp:
			ln, err = newListenerTCP(port)

			if err != nil {
				t.Errorf("Err creating TCP listener", err)
			}
			defer ln.(net.Listener).Close()
		}

		// Now try to use the same port
		_, err = Verify(test.netProto, port)
		if err == nil {
			t.Errorf("Failed to detect port was taken.")
		}

		_, err = VerifyHostPort(test.netProto, JoinHostPort("127.0.0.1", port))
		if err == nil {
			t.Errorf("Failed to detect port was taken.")
		}

	}
}
