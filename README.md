# mongo-driver-bench
Benchmarks of MongoDB Go drivers

Currently comparing mgo to the new mongo-go-driver.

To run benchmarks:
```$xslt
darmstrong@Immigrant mongo-driver-bench (master)$ go test github.com/dunkyboy/mongo-driver-bench -run=XXX -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/dunkyboy/mongo-driver-bench
BenchmarkMgoMarshalSmallStruct-8              	 2000000	       589 ns/op	     528 B/op	       3 allocs/op
BenchmarkMgoMarshalSmallNestedStruct-8        	  200000	      6718 ns/op	    2288 B/op	       5 allocs/op
BenchmarkMgoMarshalLargeStruct-8              	  300000	      4797 ns/op	    8672 B/op	      16 allocs/op
BenchmarkMgoMarshalLargeNestedStruct-8        	   50000	     37093 ns/op	   31328 B/op	      20 allocs/op
BenchmarkMgoUnmarshalSmallStruct-8            	 2000000	       930 ns/op	     280 B/op	       9 allocs/op
BenchmarkMgoUnmarshalSmallNestedStruct-8      	  200000	     10266 ns/op	     928 B/op	     108 allocs/op
BenchmarkMgoUnmarshalLargeStruct-8            	  200000	      7650 ns/op	    2880 B/op	      90 allocs/op
BenchmarkMgoUnmarshalLargeNestedStruct-8      	   20000	     69962 ns/op	    8786 B/op	     754 allocs/op
BenchmarkDriverMarshalSmallStruct-8           	 1000000	      1738 ns/op	    1416 B/op	      29 allocs/op
BenchmarkDriverMarshalSmallNestedStruct-8     	   50000	     25577 ns/op	   11232 B/op	     416 allocs/op
BenchmarkDriverMarshalLargeStruct-8           	  100000	     14864 ns/op	   12448 B/op	     218 allocs/op
BenchmarkDriverMarshalLargeNestedStruct-8     	   10000	    176309 ns/op	   70488 B/op	    2295 allocs/op
BenchmarkDriverUnmarshalSmallStruct-8         	  300000	      4136 ns/op	    1744 B/op	      65 allocs/op
BenchmarkDriverUnmarshalSmallNestedStruct-8   	   10000	    107357 ns/op	   55081 B/op	    1910 allocs/op
BenchmarkDriverUnmarshalLargeStruct-8         	    5000	    272784 ns/op	   49674 B/op	    3818 allocs/op
BenchmarkDriverUnmarshalLargeNestedStruct-8   	     500	   3051405 ns/op	  709692 B/op	   45207 allocs/op
PASS
ok  	github.com/dunkyboy/mongo-driver-bench	29.037s
```