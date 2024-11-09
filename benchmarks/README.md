## Benchmarking Results
Running the benchmark should obtain similar results regardless of the hardware, but for the sake of clarity

```
goos: darwin
goarch: amd64
pkg: github/lwileczek/goBenchmarkSerialization/benchmarks
cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
```

### Results
I don't think I'm doing FBE correctly. Any of the protobuff like things which make have a seperate
constructor have magic in them that I did not inspect. There is a chance FBE Unmarshal is hella fast
but also there is a chance it's just reading the same (last encoded value) out repeatedly

If using JSON, [sonic](https://github.com/bytedance/sonic) does appear to provide a huge performance boost,
as they suggest.

#### Marshal
| Test | Times Run | Time per Operation | Memory Used | Allocations made |
|:--|---:|---:|---:|---:|
|BenchmarkGOB/Marshal-8	              | 425,778	     | 2,928 ns/op	    |1,183 B/op	      |14 allocs/op|
|BenchmarkJSON/Marshal-8          	  | 244,279	     | 5,446 ns/op	    |3,219 B/op	      |15 allocs/op|
|BenchmarkSonicJSON/Marshal-8     	  | 689,878	     | 2,238 ns/op	    |2,060 B/op	      | 3 allocs/op|
|BenchmarkMsgPack/Marshal-8       	  | 181,344	     | 6,680 ns/op	    |3,043 B/op	      |14 allocs/op|
|BenchmarkProtobuf/Marshal-8      	  | 215,881	     | 5,324 ns/op	    |  876 B/op	      |30 allocs/op|
|BenchmarkFBE/Marshal-8           	  | 205,209	     | 5,028 ns/op	    |4,187 B/op	      |13 allocs/op|

#### Unmarshal
| Test | Times Run | Time per Operation | Memory Used | Allocations made |
|:--|---:|---:|---:|---:|
|BenchmarkGOB/Unmarshal-8         	  | 311,762	     | 3,361 ns/op	    |1,733 B/op	      |20 allocs/op|
|BenchmarkJSON/Unmarshal-8        	  |  55,278	     |20,527 ns/op	    |2,149 B/op	      |28 allocs/op|
|BenchmarkSonicJSON/Unmarshal-8   	  | 274,930	     | 3,900 ns/op	    |2,715 B/op	      |11 allocs/op|
|BenchmarkMsgPack/Unmarshal-8     	  | 107,505	     |10,412 ns/op	    |1,988 B/op	      |32 allocs/op|
|BenchmarkProtobuf/Unmarshal-8    	  | 290,875	     | 4,467 ns/op	    |2,403 B/op	      |43 allocs/op|
|BenchmarkFBE/Unmarshal-8         	  | 513,672	     | 2,149 ns/op	    |2,024 B/op	      |25 allocs/op|
