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
|BenchmarkGOBMarshal         | 206593 |   6892 ns/op |  4272 B/op  |  67 allocs/op |
|BenchmarkJSONMarshal        | 220186 |   4988 ns/op |  1632 B/op  |  12 allocs/op |
|BenchmarkMsgPackMarshal     | 478756 |  11366 ns/op |  5717 B/op  |  31 allocs/op |
|BenchmarkProtobufMarshal    | 749360 |   3581 ns/op |  1136 B/op  |  43 allocs/op |

#### Unmarshal
| Test | Times Run | Time per Operation | Memory Used | Allocations made |
|:--|---:|---:|---:|---:|
|BenchmarkGOBUnmarshal       |  52029 |  22388 ns/op | 10965 B/op  | 284 allocs/op |
|BenchmarkJSONUnmarshal      |  35520 |  30646 ns/op |  6576 B/op  | 236 allocs/op |
|BenchmarkMsgPackUnmarshal   | 125197 |  15380 ns/op |  2968 B/op  |  32 allocs/op |
|BenchmarkProtobufUnmarshal  | 391359 |   3529 ns/op |  1843 B/op  |  55 allocs/op |
