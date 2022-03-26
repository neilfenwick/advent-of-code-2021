# Day 4: High-Entropy Passphrases

## Part 1

This problem initially looks like it consists of a few simple parts:

- iterate a text file and read line-by-line
- tokenize each line (a passphrase) and split on whitespace to find words
- for each passphrase, hash each word and count unique hashes

In most higher-order languages, a dictionary / map construct will provide
the hashing and counting ability. These data structures usually hash the key
(each word of the passphrase will be a key) and will have some means of testing if a key already exists.

### Part 1 Benchmarks

```text
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day4/pwd -bench ^BenchmarkIsUniqueWords$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day4/pwd
BenchmarkIsUniqueWords-4        2000000           964 ns/op         499 B/op           2 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day4/pwd    2.897s
Success: Benchmarks passed.
```

## Part 2

The second part follows pretty much the same approach as the first, except that the test for uniqueness now includes all anagrams of each word

I took a naive approach and just sorted the characters of each word into alphabetical order, and then effectively fed that back into my Part 1 solution.  This could be optimised, but it was the simplest thing that works as a starting point for a solution.

I also introduced a stateful struct.  This isn't strictly necessary, as everything is essentially just a function call at present.  In theory a struct can be extended to hold state for the caller to interrogate about the reason(s) for the most recent result.

### Part 2 Benchmarks

```text
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day4/pwd -bench ^BenchmarkIsUniqueAnagram$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day4/pwd
BenchmarkIsUniqueAnagram-4         300000          4425 ns/op        1851 B/op          45 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day4/pwd    1.375s
Success: Benchmarks passed.
```

Part 2 could probably be improved by doing a more efficient sort of the runes in each word