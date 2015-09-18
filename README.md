## PortUtil
Quickly verify if a port is available or get a unique port (TCP/UDP).

## Usage
### Verify, VerifyTCP, VerifyUDP:
```
port, err := portutil.Verify("tcp", 8080)
if err != nil {
	log.Fatal(err)
}
```
```
port, err := portutil.VerifyUDP(8080)
if err != nil {
	log.Fatal(err)
}
```
```
port, err := portutil.VerifyTCP(8080)
if err != nil {
	log.Fatal(err)
}
```

Verify Port from a HostPort string (TCP Only)
```
serviceHost, err := portutil.VerifyHostPortTCP("127.0.0.1:8080")
if err != nil {
	log.Fatal(err)
}
```

### GetUnique, GetUniqueTCP, GetUniqueUDP:
```
port, err := portutil.GetUnique("tcp")
if err != nil {
	log.Fatal(err)
}
```
```
port, err := portutil.GetUniqueTCP
if err != nil {
	log.Fatal(err)
}
```
```
port, err := portutil.GetUniqueUDP
if err != nil {
	log.Fatal(err)
}
```