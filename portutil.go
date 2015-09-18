package portutil

import (
	"net"
	"strconv"
)

func Verify(port int) (verifiedPort int, err error) {
	address := JoinHostPort("127.0.0.1", port)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return 0, err
	}
	defer ln.Close()

	return port, nil
}

func VerifyHostPort(address string) (verifiedAddress string, err error) {
	port, err := GetPortFromAddr(address)
	if err != nil {
		return "", err
	}

	_, err = Verify(port)
	if err != nil {
		return "", err
	}

	return address, nil
}

func GetUnique() (port int, err error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
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

func GetPortFromAddr(address string) (port int, err error) {
	_, portStr, err := net.SplitHostPort(address)
	if err != nil {
		return 0, err
	}

	port, err = strconv.Atoi(portStr)
	if err != nil {
		return 0, err
	}

	return port, nil
}

func JoinHostPort(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}
