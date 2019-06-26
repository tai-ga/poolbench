# sync.Pool benchmark

## result
```
go version go1.12.6 darwin/amd64
> go test -benchmem -bench .
Name: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
PhysicalCores: 4
ThreadsPerCore: 2
LogicalCores: 8
Family 6 Model: 142
Features: CMOV,NX,MMX,MMXEXT,SSE,SSE2,SSE3,SSSE3,SSE4.1,SSE4.2,AVX,AVX2,FMA3,F16C,BMI1,BMI2,LZCNT,POPCNT,AESNI,CLMUL,HTT,RDRAND,RDSEED,ADX,MPX,ERMS,RDTSCP,CX16,SGX,IBPB,STIBP
Cacheline bytes: 64
L1 Data Cache: 32768 bytes
L1 Instruction Cache: 32768 bytes
L2 Cache: 262144 bytes
L3 Cache: 8388608 bytes
We have Streaming SIMD Extensions
---
goos: darwin
goarch: amd64
pkg: github.com/tai-ga/poolbench
BenchmarkWithoutPool/________1-8               200000000                 9.33 ns/op               1 B/op          1 allocs/op
BenchmarkWithoutPool/_______10-8               100000000                15.9  ns/op              16 B/op          1 allocs/op
BenchmarkWithoutPool/______100-8                50000000                26.0  ns/op             112 B/op          1 allocs/op
BenchmarkWithoutPool/_____1000-8                20000000               111    ns/op            1024 B/op          1 allocs/op
BenchmarkWithoutPool/____10000-8                 2000000               601    ns/op           10240 B/op          1 allocs/op
BenchmarkWithoutPool/___100000-8                  300000              5452    ns/op          106496 B/op          1 allocs/op
BenchmarkWithoutPool/__1000000-8                   30000             48049    ns/op         1007624 B/op          1 allocs/op
BenchmarkWithoutPool/_10000000-8                    3000            463783    ns/op        10002443 B/op          1 allocs/op
BenchmarkWithoutPool/100000000-8                     500           2448606    ns/op       100007954 B/op          1 allocs/op
BenchmarkWithPool/________1-8                   30000000                41.6  ns/op              32 B/op          1 allocs/op
BenchmarkWithPool/_______10-8                   30000000                41.4  ns/op              32 B/op          1 allocs/op
BenchmarkWithPool/______100-8                   30000000                41.7  ns/op              32 B/op          1 allocs/op
BenchmarkWithPool/_____1000-8                   30000000                41.8  ns/op              32 B/op          1 allocs/op
BenchmarkWithPool/____10000-8                   30000000                41.7  ns/op              32 B/op          1 allocs/op
BenchmarkWithPool/___100000-8                   30000000                42.0  ns/op              32 B/op          1 allocs/op
BenchmarkWithPool/__1000000-8                   30000000                43.0  ns/op              38 B/op          1 allocs/op
BenchmarkWithPool/_10000000-8                   30000000                43.6  ns/op              58 B/op          1 allocs/op
BenchmarkWithPool/100000000-8                   30000000                45.5  ns/op              62 B/op          1 allocs/op
PASS
ok      github.com/tai-ga/poolbench     28.374s
```
