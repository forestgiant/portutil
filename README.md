## PortUtil
Quickly verify if a port is available. 
This is useful if your microservice architecture requires a service to always be on the same port.

## Usage
Verify Port (TCP Only)
```
port, err := portutil.Verify(8080)
if err != nil {
	log.Fatal(err)
}
```

Verify Port from a HostPort string (TCP Only)
```
serviceHost, err := portutil.VerifyHostPort("127.0.0.1:8080")
if err != nil {
	log.Fatal(err)
}
```

Get a unique port (TCP Only)
```
port, err := portutil.GetUnique()
if err != nil {
	log.Fatal(err)
}
```

Verify Port by Network Protocol. "udp" or "tcp"
```
port, err := portutil.VerifyByNet("udp", 8053)
if err != nil {
	log.Fatal(err)
}
```