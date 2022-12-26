
Start server `go run server.go 9800`
Start Client 

You can write to a connection and read data out of a connection
with the following
```go
_, err = conn.Write([]byte("This is a message"))
if err != nil {
        println("Write data failed:", err.Error())
            os.Exit(1)
}

// buffer to get data
received := make([]byte, 1024)
_, err = conn.Read(received)
if err != nil {
        println("Read data failed:", err.Error())
            os.Exit(1)
}
````


## Custom TCP server with Serialization
https://appliedgo.net/networking/
