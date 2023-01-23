## Benchmarking Results
Running the benchmark should obtain similar results regardless of the hardware, but for the sake of clarity

```
goos: darwin
goarch: amd64
pkg: github/lwileczek/goBenchmarkSerialization/benchmarks
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
```

### Results
GOBs seem to perform the worst with Protobufs performing the best.
Although MesgPack uses a lot of memory and is slow to marshal, it is much faster unmarshalling compared to JSON.
It's possible if I tried harder on the serializer size I could get it to perform better there.

#### Marshal
| Test | Times Run | Time per Operation | Memory Used | Allocations made |
|:--|---:|---:|---:|---:|
|BenchmarkGOBMarshal         | 206,593 |   6,892 ns/op |  4,272 B/op  |  67 allocs/op |
|BenchmarkJSONMarshal        | 220,186 |   4,988 ns/op |  1,632 B/op  |  12 allocs/op |
|BenchmarkMsgPackMarshal     | 478,756 |  11,366 ns/op |  5,717 B/op  |  31 allocs/op |
|BenchmarkProtobufMarshal    | 749,360 |   3,581 ns/op |  1,136 B/op  |  43 allocs/op |

#### Unmarshal
| Test | Times Run | Time per Operation | Memory Used | Allocations made |
|:--|---:|---:|---:|---:|
|BenchmarkGOBUnmarshal       |  52,029 |  22,388 ns/op | 10,965 B/op  | 284 allocs/op |
|BenchmarkJSONUnmarshal      |  35,520 |  30,646 ns/op |  6,576 B/op  | 236 allocs/op |
|BenchmarkMsgPackUnmarshal   | 125,197 |  15,380 ns/op |  2,968 B/op  |  32 allocs/op |
|BenchmarkProtobufUnmarshal  | 391,359 |   3,529 ns/op |  1,843 B/op  |  55 allocs/op |
