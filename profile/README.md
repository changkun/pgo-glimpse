```bash
$ go test -v -run=none -bench=. -count=10 --cpuprofile cpu.pprof | tee bench-mutex.txt
goos: darwin
goarch: arm64
pkg: changkun.de/x/pgo-glimpse/profile
BenchmarkCounter
BenchmarkCounter-8      10036978               119.1 ns/op
BenchmarkCounter-8      10003890               117.2 ns/op
BenchmarkCounter-8      10037259               110.6 ns/op
BenchmarkCounter-8      12362572               119.2 ns/op
BenchmarkCounter-8       9567883               122.3 ns/op
BenchmarkCounter-8      10029978               119.6 ns/op
BenchmarkCounter-8      10369540               119.8 ns/op
BenchmarkCounter-8       9718051               119.7 ns/op
BenchmarkCounter-8       9915026               118.2 ns/op
BenchmarkCounter-8      10082162               118.7 ns/op
PASS
ok      changkun.de/x/pgo-glimpse/profile       13.513s

$ go test -v -run=none -bench=. -count=10 --cpuprofile cpu2.pprof | tee bench-automic.txt
goos: darwin
goarch: arm64
pkg: changkun.de/x/pgo-glimpse/profile
BenchmarkCounter
BenchmarkCounter-8      22162659                53.77 ns/op
BenchmarkCounter-8      19759862                58.41 ns/op
BenchmarkCounter-8      25623642                51.70 ns/op
BenchmarkCounter-8      25068218                58.81 ns/op
BenchmarkCounter-8      22502162                58.07 ns/op
BenchmarkCounter-8      23885677                49.61 ns/op
BenchmarkCounter-8      20504496                61.07 ns/op
BenchmarkCounter-8      27760854                50.10 ns/op
BenchmarkCounter-8      22552168                50.04 ns/op
BenchmarkCounter-8      26357139                54.19 ns/op
PASS
ok      changkun.de/x/pgo-glimpse/profile       15.947s

$ benchstat bench-mutex.txt bench-automic.txt
name       old time/op  new time/op  delta
Counter-8   119ns ± 1%    55ns ±12%  -54.11%  (p=0.000 n=8+10)
```

```
$ go tool pprof cpu.pprof
$ go tool pprof cpu2.pprof
$ go tool pprof --diff_base cpu.pprof cpu2.pprof
$ go test -v -run=none -bench=. -benchtime=10s --cpuprofile cpu2.pprof | tee bench-automic.txt
$ go test -v -run=none -bench=. -benchtime=10s --cpuprofile cpu.pprof | tee bench-mutex.txt
$ go tool pprof --diff_base cpu.pprof cpu2.pprof
```