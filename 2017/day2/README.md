# Day 2: Corruption Checksum

## Initial impressions

Straight away this looks like a streaming problem, feeding into an accumulator function chain: map-reduce.  One function to calculate the row checksum, and one function to aggregate the row checksums for the file checksum.  In C# this could feed into a LINQ call to `Select()`, followed by `Aggregate()`.

We could get away with assuming the buffer was in memory in the Day 1 problem (since it is a circular buffer...).  The Day 2 challenge could be made to be a little more reality-proof if we don't try to load the entire input file into memory, but rather stream it.  I don't want to get picked on by [Ayende](https://ayende.com) for trying to `File.ReadAll()` a hypothetical massive file that wouldn't fit into memory :)

As a stretch, the different operations could potentially be split using Go's channels as the communication mechanism between them, allowing some asynchronous processing.  That would probably be over-engineered though because the 'bottleneck' is likely to be the part that is reading the input stream.

### Processing the input

I tried to find the idiomatic Go way of reading a file of integers and discovered the `bufio.Scanner`.  At first, I thought I would try to implement the `Scanner` interface, only to find there isn't one.  It probably would not have helped much anyway because the `.Text()` func would have returned a `string` and I really wanted to return `[]int`.  That left me with decorating the `Scanner` and creating my own type to wrap it, whilst parsing the input lines as `[]int`.

## Part 2 reflections

I had to resist my object-oriented C# background here.  When part 2 of the problem introduced a different calculation method for the checksum, my first instinct was to create a new *type*.  In C# that would have been a new *class*, and there was a temptation to clone & rename `Checksum`; to create `DiffChecksum` and `DivideChecksum` structs, each with their own function implementation - **the Strategy Pattern**.  But in Go, *functions can be types*, so the Strategy need not be implemented with classes & structs.

## Benchmarks

```text
Part 1:
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day2 -bench ^BenchmarkDiffChecksum$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day2
BenchmarkDiffChecksum-4       300000000             5.95 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day2    2.390s
Success: Benchmarks passed.

Part 2:
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day2 -bench ^BenchmarkModulusChecksum$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day2
BenchmarkModulusChecksum-4       200000000             7.61 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day2    2.347s
Success: Benchmarks passed.
```