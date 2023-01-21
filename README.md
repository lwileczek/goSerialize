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
But Deku appears to only be for Rust, much like GOB for Go, and Pickle for Python.

The Repository for Fast Binary Encoding (FBE) wrote their own benchmark tests within their repository.
Here are their tests written in Go: https://github.com/chronoxor/FastBinaryEncoding/tree/master/projects/Go/benchmarks

[Alec Thomas: Serialization Benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)


### Why I wrote this
I recently saw the benchmarking video by Primeagen, mentioned above, which reminded me about MessagePack.
I heard of MsgPack years ago in this [PyCon 2014 Vid](https://www.youtube.com/watch?v=7KnfGDajDQw) about avoiding Python Pickles.
However, It seems like Protobufs are far more popular today. 
I'd love to see how easy it is to use them and how big of a performance improvement we see. 

#### Why each encoding was picked
| Encoding | Reason |
|:---|:---|
| JSON | The control / base case |
| BSON | Seems like possibly an easy speed boost without having to try too much which may be worth it in its self |
| Protobuf | Popular |
| GOB | Supposed to be easy within Go |
| FBE | makes some __bold__ claims but doesn't seem popular so I am sceptical yet want to see the results |
| MsgPack | Supposed to be easy to work with, fast, and semi-legable |

#### Runner up types
Another set of runner up data formats:
 - ion
 - xml
 - Apache Thrift (No Go?)

## Running the Benchmark
Start the server and run
```bash
go run main.go -s 30
```

## Results

|format | x | y | difficulty |
|:---|---:|---:|:---:|
|json | 1 | 2| :+1: |
|bson | 1 | 2| :grin: |
|protobufs| 1 | 2| :grimacing: |
|msgpack| 1 | 2| :smirk: |
|FastBinaryEncoding| 1 | 2| :sweat: |
|gob| 1 | 2| :grin: |

The emoji's are my subjective rating for how difficult each encoding was to work with.

### Poking at binary data
An example of each data type has been saved to file under the `data` directory.
You can poke at the binary files with `hexdump -C example.<format> | less`
To see how much of the data can be picked out of the binary files. 
For JSON, you can just open the file since it's just a string.
