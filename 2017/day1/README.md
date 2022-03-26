# Day 1: Inverse Captcha

## Reflections

### Part 1

This could easily be a simple coding kata exercise for TDD. Was quite straightforward to build up from the simple case in the [Day 1 examples](https://adventofcode.com/2017/day/1).

### Part 2

After learning from the approach in Pt2, where I used modulus to implement the circular buffer, I changed my approach in pt1 from a look-behind (previous) to a look-ahead (next)... what I should have done from the start, really.  The look-behind approach would probably apply to a streamed data-source, because all you would have access to is the current and previous values.  In this case, we can "cheat" because the problem states a circular buffer, which implies that all the data is "addressable", in this case, using array indices.

## Benchmarks

The `BenchmarkSumConsecutiveIntegers` applies to my second, look-ahead, implementation with the buffer all in memory

```text
# Day 1 Pt1
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day1/calc -bench ^BenchmarkSumConsecutiveIntegers$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day1/calc
BenchmarkSumConsecutiveIntegers-4        1000000          1448 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day1/calc    1.466s
Success: Benchmarks passed.

# Day 1 Pt2
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day1/calc -bench ^BenchmarkSumOppositeIntegers$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day1/calc
BenchmarkSumOppositeIntegers-4         100000         17743 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day1/calc    1.960s
Success: Benchmarks passed.
```