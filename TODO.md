
 - [ ] Hopefully avoid orphan go routines
 - [ ] use persistent data until the connection is closed by client
 - [ ] reduce repeated code in both ./server and ./client. gob, json, and msgpack are almost identical functions so we should be able to do something to reduce the repeated code
 - [ ] Homogonize how we send and receive. Protocol buffers are not using buffered reads/writes

