# Status
Alpha

# Benchmark Serialization
This is part of two projects to test the serialization and transfer
of data between a client and server using the following formats:

 - [protobufs](https://developers.google.com/protocol-buffers/)
 - json
 - [bson](https://bsonspec.org/)
 - [msgpack](https://msgpack.org/)
 - [FastBinaryEncoding (FBE)](https://github.com/chronoxor/FastBinaryEncoding)

### Other Benchmarks
I liked this benchmark of Deku vs. JSON by Primeagen: [YouTube Vid](https://www.youtube.com/watch?v=MuCK81q1edU)
But Deku appears to only be for Rust. 
Let's see if we can create a similar benchmarking test (in Go :heart:).

 - FBE Go code: https://github.com/chronoxor/FastBinaryEncoding/tree/master/projects/Go/benchmarks

### Why I wrote this
I heard of MsgPack years ago in this [PyCon 2014 Vid](https://www.youtube.com/watch?v=7KnfGDajDQw) about avoiding Python Pickles.
However, It seems like Protobufs are far more popular today. 
I'd love to see how easy it is to use both of them. 
BSON seems like possibly an easy speed boost without having to try too much which may be worth it in its self.
FBE makes some __bold__ claims but doesn't seem popular so I am sceptical yet want to see the results.

#### Runner up types
Another set of runner up data formats:
 - ion
 - xml (boo)
 - Apache Thrift (No Go?)

## Running the Benchmark
Start the server and run
```bash
go run main.go -s 30
```

## Results

|format | x | y
|:---|---:|---:|
|json | 1 | 2|
|bson | 1 | 2|
|protobufs| 1 | 2|
|msgpack| 1 | 2|
|FastBinaryEncoding| 1 | 2|


## ToDo
 - More on error groups: https://bostonc.dev/blog/go-errgroup
 - Write classic go Benchmark code like: https://github.com/chronoxor/FastBinaryEncoding/tree/master/projects/Go/benchmarks
 - Read about Gobs of data: https://go.dev/blog/gob Looks like a language specific encoding like a Python Pickle
