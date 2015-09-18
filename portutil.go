package portutil

import (
	"net"
	"strconv"
)

func Verify(port int) (verifiedPort int, err error) {
	addr := JoinHostPort("127.0.0.1", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer ln.Close()

	return port, nil
}

func VerifyByNet(netProto string, port int) (verifiedPort int, err error) {

	switch netProto {
	case "udp":
		addr := JoinHostPort("127.0.0.1", port)
		udpAddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			return 0, err
		}

		ln, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			return 0, err
		}
		defer ln.Close()
	case "tcp":
		_, err = Verify(port)
		if err != nil {
			return 0, err
		}
	}

	return port, nil
}

func VerifyHostPort(addr string) (verifiedAddr string, err error) {
	port, err := GetPortFromAddr(addr)
	if err != nil {
		return "", err
	}

	_, err = Verify(port)
	if err != nil {
		return "", err
	}

	return addr, nil
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

func JoinHostPort(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}
