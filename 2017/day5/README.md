# Day 5: A Maze of Twisty Trampolines, All Alike

Initially this looks like a fairly straight-forward and typical TDD Kata exercise.

A list of size `1`, any non-zero value will cause an offset outside the list - exit with success
Any list greater than one, add the value to the current offset, increment the current value, jump to new offset, repeat.

## Part 1

### Example input

```text
0
3
0
1
-3
```

Expected: `5` jumps to reach the end of the list

#### Benchmarks

```text
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day5/jmp -bench ^BenchmarkDefaultJumpStrategy$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day5/jmp
BenchmarkDefaultJumpStrategy-4           2000       1036183 ns/op        9472 B/op           1 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day5/jmp    2.182s
Success: Benchmarks passed.
```

## Part 2

Another example of the strategy pattern.  Introduce a new function to alter the jump offset calculation, based on a different set of criteria

### Example input - part 2

```text
0
3
0
1
-3
```

Expected: `10` jumps to reach the end of the list

#### Benchmarks - part 2

```text
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day5/jmp -bench ^BenchmarkNewJumpStrategy$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day5/jmp
BenchmarkNewJumpStrategy-4             10     141762771 ns/op        9472 B/op           1 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day5/jmp    1.563s
Success: Benchmarks passed.
```