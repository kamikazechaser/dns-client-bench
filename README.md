## dns-client-bench

### Results

Scope: Performing an `A` lookup (most common use case).

When resolving through DoT and TCP,
[`phuslu/fastdns`](https://github.com/phuslu/fastdns) is between 2x and 4x
faster than both [`miekg/dns`](github.com/miekg/dns) and `stdlib`. When it comes
to UDP resolutions, all packages are more or less equivalent with
`phuslu/fastdns` being fractionally faster. In all cases, `miekg/dns` is
fractionally faster than `stdlib`.

When it comes memory usage, `phuslu/fastdns` is extremely efficient in all cases
while `miekg/dns` and `std` have comparable usage.

```bash
goos: linux
goarch: amd64
pkg: github.com/kamikazechaser/dns-client-bench
cpu: Intel Xeon Processor (Skylake, IBRS)
BenchmarkDOTSTDLibDNS
BenchmarkDOTSTDLibDNS-2            	     705	   8303075 ns/op	   80203 B/op	     702 allocs/op
BenchmarkDOTSTDLibDNS-2            	     716	   8258800 ns/op	   80222 B/op	     702 allocs/op
BenchmarkDOTSTDLibDNS-2            	     714	   8338623 ns/op	   80256 B/op	     702 allocs/op
BenchmarkDOTPhusluFastDNS
BenchmarkDOTPhusluFastDNS-2        	    2467	   2373994 ns/op	     561 B/op	       7 allocs/op
BenchmarkDOTPhusluFastDNS-2        	    2517	   2414586 ns/op	     550 B/op	       6 allocs/op
BenchmarkDOTPhusluFastDNS-2        	    2454	   2420993 ns/op	     563 B/op	       7 allocs/op
BenchmarkDOTMiekgDNS
BenchmarkDOTMiekgDNS-2             	     746	   8019842 ns/op	   78867 B/op	     793 allocs/op
BenchmarkDOTMiekgDNS-2             	     747	   7916746 ns/op	   78865 B/op	     793 allocs/op
BenchmarkDOTMiekgDNS-2             	     742	   8009323 ns/op	   78867 B/op	     793 allocs/op
BenchmarkSTDLibDNSHostResolver
BenchmarkSTDLibDNSHostResolver-2   	   36987	    148781 ns/op	    4344 B/op	      49 allocs/op
BenchmarkSTDLibDNSHostResolver-2   	   48937	    144887 ns/op	    4344 B/op	      49 allocs/op
BenchmarkSTDLibDNSHostResolver-2   	   44881	    146155 ns/op	    4344 B/op	      49 allocs/op
BenchmarkSTDLibDNS
BenchmarkSTDLibDNS-2               	    1262	   4765375 ns/op	    4650 B/op	      53 allocs/op
BenchmarkSTDLibDNS-2               	    1242	   4699738 ns/op	    4650 B/op	      53 allocs/op
BenchmarkSTDLibDNS-2               	    1291	   4714137 ns/op	    4649 B/op	      53 allocs/op
BenchmarkPhusluFastDNS
BenchmarkPhusluFastDNS-2           	    2284	   2524174 ns/op	     820 B/op	      14 allocs/op
BenchmarkPhusluFastDNS-2           	    2424	   2444125 ns/op	     820 B/op	      14 allocs/op
BenchmarkPhusluFastDNS-2           	    2361	   2482972 ns/op	     820 B/op	      14 allocs/op
BenchmarkMiekgDNS
BenchmarkMiekgDNS-2                	    1294	   4593856 ns/op	    2050 B/op	      36 allocs/op
BenchmarkMiekgDNS-2                	    1298	   4666096 ns/op	    2051 B/op	      36 allocs/op
BenchmarkMiekgDNS-2                	    1291	   4594621 ns/op	    2048 B/op	      36 allocs/op
BenchmarkUDPSTDLibDNS
BenchmarkUDPSTDLibDNS-2            	    2266	   2717213 ns/op	    4389 B/op	      49 allocs/op
BenchmarkUDPSTDLibDNS-2            	    2190	   2671163 ns/op	    4422 B/op	      49 allocs/op
BenchmarkUDPSTDLibDNS-2            	    2331	   2633149 ns/op	    4420 B/op	      49 allocs/op
BenchmarkUDPPhusluFastDNS
BenchmarkUDPPhusluFastDNS-2        	    2190	   2442131 ns/op	     377 B/op	      10 allocs/op
BenchmarkUDPPhusluFastDNS-2        	    2233	   2448309 ns/op	     370 B/op	      10 allocs/op
BenchmarkUDPPhusluFastDNS-2        	    2205	   2436252 ns/op	     375 B/op	      10 allocs/op
BenchmarkUDPMiekgDNS
BenchmarkUDPMiekgDNS-2             	    2450	   2489418 ns/op	    2139 B/op	      29 allocs/op
BenchmarkUDPMiekgDNS-2             	    2364	   2485459 ns/op	    2139 B/op	      29 allocs/op
BenchmarkUDPMiekgDNS-2             	    2299	   2509677 ns/op	    2139 B/op	      29 allocs/op
```
