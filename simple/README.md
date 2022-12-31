# Simple TCP Server
A simple example of a custom TCP server sending binary encoded data back and forth.
Data is encoded using Go [Gobs](https://blog.golang.org/gobs-of-data) and sent
to a custom server, which reads it, encodes it in JSON to pretty print, and returns
a string "Successful transaction".

## Why
This demonstrates how to send data between two services via bytes over TCP.

## How to Use
Start server `go run main.go client.go server.go`

The Loop runs five times and then waits for you to hit a key before exiting.
