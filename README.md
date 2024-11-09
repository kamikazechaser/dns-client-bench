### dns-client-bench

```bash
go test -v -run=none -benchmem -bench=.
goos: linux
goarch: amd64
pkg: github.com/kamikazechaser/dns-client-bench
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkDOTSTDLibDNS
BenchmarkDOTSTDLibDNS-12                     207           7759207 ns/op           81131 B/op        707 allocs/op
BenchmarkDOTPhusluFastDNS
BenchmarkDOTPhusluFastDNS-12                 888           1306652 ns/op            8323 B/op         81 allocs/op
BenchmarkDOTMiekgDNS
BenchmarkDOTMiekgDNS-12                      492           4057353 ns/op           79368 B/op        792 allocs/op
------------------------------------------------------------------------------------------------------------------
BenchmarkSTDLibDNS
BenchmarkSTDLibDNS-12                      25766             44565 ns/op            4344 B/op         49 allocs/op
BenchmarkPhusluFastDNS
BenchmarkPhusluFastDNS-12                   3164            359611 ns/op             120 B/op          3 allocs/op
BenchmarkMiekgDNS
BenchmarkMiekgDNS-12                        3198            641758 ns/op            2144 B/op         30 allocs/op
PASS
ok      github.com/kamikazechaser/dns-client-bench      10.916s
```
