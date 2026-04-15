# crypto/crand benchmarks

Run the benchmark:

```text
make bench name=crypto-crand
```

Go 1.26.1:

```text
goos: darwin
goarch: arm64
pkg: solod.dev/bench/crypto-crand
cpu: Apple M1
Benchmark_Read_4-8     17044155      69.31 ns/op      57.71 MB/s
Benchmark_Read_32-8     4952852     242.2 ns/op      132.11 MB/s
Benchmark_Read_4K-8      996122    1215 ns/op       3370.07 MB/s
Benchmark_Text-8        4527999     263.9 ns/op       98.53 MB/s
```

So:

```text
Benchmark_Read_4       29645734      39.75 ns/op     100.63 MB/s
Benchmark_Read_32       5685102     210.6 ns/op      151.96 MB/s
Benchmark_Read_4K       1000000    1184 ns/op       3459.46 MB/s
Benchmark_Text          5626722     212.8 ns/op      122.19 MB/s
```
