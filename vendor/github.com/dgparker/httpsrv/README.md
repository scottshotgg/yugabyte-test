# httpsrv
```go get github.com/dgparker/httpsrv```

Common overrides and utilities for the go standard library http.Server

## Usage
```
// initialize your handler which implements http.Handler
handler := newHandler()

// intialize httpsrv with default settings
// (*optional call New() and your own configuration for more control)
server := httpsrv.NewWithDefault(handler)

// set server port
server.Addr = ":9001"

// start the server
log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
```