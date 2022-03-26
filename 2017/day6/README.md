# Day 6: Memory Reallocation

This problem initially looks like an iterative rebalancing algorithm, combined with some form of "hashing" function that is performed after each pass.  We stop iterating when we encounter a "hash" that we have observed before.

## Benchmarks - Part 1

```text
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day6/mem -bench ^BenchmarkRebalance$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day6/mem
BenchmarkRebalance-4            200       8234354 ns/op     2313436 B/op      141888 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day6/mem    2.494s
Success: Benchmarks passed.
```

## Part 2

Essentially the problem was solved in such a way that the output from Part 1 could be fed back into the calculation to produce the result for Part 2