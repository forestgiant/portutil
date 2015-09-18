package portutil

import (
	"net"
	"testing"
)

func TestVerify(t *testing.T) {
	testPort := 8000

	port, err := Verify(testPort)
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

func TestVerifyByNet(t *testing.T) {
	testPort := 9000

	tests := []struct {
		netProto string
	}{
		{"udp"},
		{"tcp"},
	}
	for _, test := range tests {
		port, err := VerifyByNet(test.netProto, testPort)
		if err != nil {
			t.Errorf("Couldn't Verify port on UDP: %s. Error: %s", port, err)
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
	addr, err := VerifyHostPort(testAddr)
	if err != nil {
		t.Errorf("Couldn't VerifyHostPort: %s. Error: %s", addr, err)
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
		// Verify TCP
		address := JoinHostPort("127.0.0.1", test.port)
		ln, err := net.Listen("tcp", address)
		if err != nil {
			t.Errorf("Err creating listener", err)
		}
		defer ln.Close()

		address, err = VerifyHostPort(address)
		if err == nil {
			t.Errorf("Failed to detect port was taken.")
		}

		// Verify UDP
		address = JoinHostPort("127.0.0.1", test.port)
		udpAddr, err := net.ResolveUDPAddr("udp", address)
		if err != nil {
			t.Errorf("Err ResolveUDPAddr", err)
		}

		udpLn, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			t.Errorf("Err creating UDP listener", err)
		}
		defer udpLn.Close()

		_, err = VerifyByNet("udp", test.port)
		if err == nil {
			t.Errorf("Failed to detect port was taken.")
		}

	}
}
