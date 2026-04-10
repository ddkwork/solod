# maps benchmarks

Requires GCC/Clang and mimalloc (for heap allocations in So). If mimalloc isn't available, the benchmarks will use the default libc allocator, which is much slower.

Run the benchmark:

```text
make bench name=maps
```

Go 1.26.1:

```text
goos: darwin
goarch: arm64
pkg: solod.dev/bench/maps
cpu: Apple M1

Benchmark_IntSet-8       31490    35645 ns/op     74264 B/op    20 allocs/op
Benchmark_IntPre-8      123108     9676 ns/op     36944 B/op     5 allocs/op
Benchmark_IntGet-8      216240     5594 ns/op         0 B/op     0 allocs/op
Benchmark_IntDel-8       49941    23968 ns/op     36944 B/op     5 allocs/op

Benchmark_StrSet-8       25107    47805 ns/op    108760 B/op    20 allocs/op
Benchmark_StrPre-8       81638    14699 ns/op     54608 B/op     5 allocs/op
Benchmark_StrGet-8      131050     9216 ns/op         0 B/op     0 allocs/op
Benchmark_StrDel-8       35484    33819 ns/op     54608 B/op     5 allocs/op
```

So (mimalloc):

```text
Benchmark_IntSet         41629    26333 ns/op     65472 B/op    15 allocs/op
Benchmark_IntPre        137805     8813 ns/op     49152 B/op     3 allocs/op
Benchmark_IntGet        780385     1581 ns/op         0 B/op     0 allocs/op
Benchmark_IntDel         79515    14889 ns/op     49152 B/op     3 allocs/op

Benchmark_StrSet         38630    31055 ns/op     87296 B/op    15 allocs/op
Benchmark_StrPre         99391    12101 ns/op     65536 B/op     3 allocs/op
Benchmark_StrGet        117486    10170 ns/op         0 B/op     0 allocs/op
Benchmark_StrDel         49550    24227 ns/op     65536 B/op     3 allocs/op
```

So (arena):

```text
Benchmark_IntSet         47002    25515 ns/op     65472 B/op    15 allocs/op
Benchmark_IntPre        137667     8704 ns/op     49152 B/op     3 allocs/op
Benchmark_IntGet        780284     1537 ns/op         0 B/op     0 allocs/op
Benchmark_IntDel         80742    14859 ns/op     49152 B/op     3 allocs/op

Benchmark_StrSet         39026    30749 ns/op     87296 B/op    15 allocs/op
Benchmark_StrPre         98018    12233 ns/op     65536 B/op     3 allocs/op
Benchmark_StrGet        120942     9907 ns/op         0 B/op     0 allocs/op
Benchmark_StrDel         49689    24392 ns/op     65536 B/op     3 allocs/op
```

So (built-in map):

```text
Benchmark_IntSet        390433     3109 ns/op
Benchmark_IntGet        464288     2577 ns/op

Benchmark_StrSet        180883     6585 ns/op
Benchmark_StrGet        112964    10531 ns/op
```
