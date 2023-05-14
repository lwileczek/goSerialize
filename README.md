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

# On Micro-Benchmarks
From one of the greats:

> On benchmarks
>
> I've spent a lot of time benchmarking Protobuf and Cap'n Proto. One thing I've learned in the process is that most simple benchmarks you can create will not give you realistic results.
>
> First, any serialization format (even JSON) can "win" given the right benchmark case. Different formats will perform very, very differently depending on the content. Is it string-heavy, number-heavy, or object heavy (i.e. with deep message trees)? Different formats have different strengths here (Cap'n Proto is incredibly good at numbers, for example, because it doesn't transform them at all; JSON is incredibly bad at them). Is your message size incredibly short, medium-length, or very large? Short messages will mostly exercise the setup/teardown code rather than body processing (but setup/teardown is important -- sometimes real-world use cases involve lots of small messages!). Very large messages will bust the L1/L2/L3 cache and tell you more about memory bandwidth than parsing complexity (but again, this is important -- some implementations are more cache-friendly than others).
>
> Even after considering all that, you have another problem: Running code in a loop doesn't actually tell you how it performs in the real world. When run in a tight loop, the instruction cache stays hot and all the branches become highly predictable. So a branch-heavy serialization (like protobuf) will have its branching cost swept under the rug, and a code-footprint-heavy serialization (again... like protobuf) will also get an advantage. This is why micro-benchmarks are only really useful to compare code against other versions of itself (e.g. to test minor optimizations), NOT to compare completely different codebases against each other. To find out how any of this performs in the real world, you need to measure a real-world use case end-to-end. But... to be honest, that's pretty hard. Few people have the time to build two versions of their whole app, based on two different serializations, to see which one wins...

from [Kenton Varda](https://stackoverflow.com/questions/61347404/protobuf-vs-flatbuffers-vs-capn-proto-which-is-faster#answer-61370743)
