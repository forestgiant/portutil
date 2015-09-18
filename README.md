## PortUtil
Quickly verify if a port is available. 
This is useful if your microservice architecture requires a service to always be on the same port.

## Typical Usage
```
port, err := portutil.Verify(8080)
if err != nil {
	log.Fatal(err)
}
```