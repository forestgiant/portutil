package portutil

import (
	"net"
	"strconv"
)

const (
	tcp = "tcp"
	udp = "udp"
)

// Verifies if a port is available on TCP
func VerifyTCP(port int) (verifiedPort int, err error) {
	ln, err := newListenerTCP(port)
	if err != nil {
		return 0, err
	}
	defer ln.Close()

	return port, nil
}

// Verifies if a port is available on UDP
func VerifyUDP(port int) (verifiedPort int, err error) {
	ln, err := newListenerUDP(port)
	if err != nil {
		return 0, err
	}
	defer ln.Close()

	return port, nil
}

// Verifies if a port is available on "udp" or "tcp"
func Verify(netProto string, port int) (verifiedPort int, err error) {
	switch netProto {
	case udp:
		_, err = VerifyUDP(port)
		if err != nil {
			return 0, err
		}
	case tcp:
		_, err = VerifyTCP(port)
		if err != nil {
			return 0, err
		}
	}

	return port, nil
}

// Wrapper function for VerifyTCP to easily accept address string
func VerifyHostPortTCP(addr string) (verifiedAddr string, err error) {
	port, err := GetPortFromAddr(addr)
	if err != nil {
		return "", err
	}

	_, err = VerifyTCP(port)
	if err != nil {
		return "", err
	}

	return addr, nil
}

// Get a unique port TCP
func GetUniqueTCP() (port int, err error) {
	ln, err := newListenerTCP(0)
	if err != nil {
		return 0, err
	}
	defer ln.Close()

	port, err = GetPortFromAddr(ln.Addr().String())
	if err != nil {
		return 0, err
	}

	return port, nil
}

// Get a unique port UDP
func GetUniqueUDP() (port int, err error) {
	ln, err := newListenerUDP(0)
	if err != nil {
		return 0, err
	}
	defer ln.Close()

	port, err = GetPortFromAddr(ln.LocalAddr().String())
	if err != nil {
		return 0, err
	}

	return port, nil
}

// Get a unique port on "udp" or "tcp"
func GetUnique(netProto string) (port int, err error) {
	switch netProto {
	case udp:
		port, err = GetUniqueUDP()
		if err != nil {
			return 0, err
		}
	case tcp:
		port, err = GetUniqueTCP()
		if err != nil {
			return 0, err
		}
	}

	return port, nil
}

// Helper function to quickly get the port from an addr string
func GetPortFromAddr(addr string) (port int, err error) {
	_, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return 0, err
	}

	port, err = strconv.Atoi(portStr)
	if err != nil {
		return 0, err
	}

	return port, nil
}

// Wrapper function for net.JoinHostPort to easily
// pass port as an int instead of a string
func JoinHostPort(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}

// Helper function to create a new UDP Listener
func newListenerUDP(port int) (*net.UDPConn, error) {
	addr := JoinHostPort("127.0.0.1", port)
	udpAddr, err := net.ResolveUDPAddr(udp, addr)
	if err != nil {
		return nil, err
	}

	ln, err := net.ListenUDP(udp, udpAddr)
	if err != nil {
		return nil, err
	}

	return ln, nil
}

// Helper function to create new UDP Listener
func newListenerTCP(port int) (net.Listener, error) {
	addr := JoinHostPort("127.0.0.1", port)
	ln, err := net.Listen(tcp, addr)
	if err != nil {
		return nil, err
	}

	return ln, nil
}
