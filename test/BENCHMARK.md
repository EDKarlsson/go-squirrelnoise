# Benchmarks

```shell
> go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/edkarlsson/go-squirrelnoise/test
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkSquirrelNoise5-12         	1000000000	         0.3246 ns/op
BenchmarkGet1dNoiseUint-12         	474090829	         2.496 ns/op
BenchmarkGet2dNoiseUint-12         	404122176	         3.032 ns/op
BenchmarkGet3dNoiseUint-12         	365520380	         3.264 ns/op
BenchmarkGet4dNoiseUint-12         	325718362	         3.623 ns/op
BenchmarkGet1dNoiseZeroToOne-12    	362215599	         3.368 ns/op
BenchmarkGet2dNoiseZeroToOne-12    	403692505	         3.006 ns/op
BenchmarkGet3dNoiseZeroToOne-12    	363442056	         3.361 ns/op
BenchmarkGet4dNoiseZeroToOne-12    	319253066	         3.695 ns/op
BenchmarkGet1dNoiseNegToOne-12     	325833655	         3.447 ns/op
BenchmarkGet2dNoiseNegToOne-12     	397728822	         3.040 ns/op
BenchmarkGet3dNoiseNegToOne-12     	362868169	         3.323 ns/op
BenchmarkGet4dNoiseNegToOne-12     	325722782	         3.626 ns/op
PASS
ok  	github.com/edkarlsson/go-squirrelnoise/test	19.273s
```