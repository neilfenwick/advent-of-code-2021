# Go TDD through Advent Of Code 2017

This repository contains solutions for experiments in Go, using TDD, and driven using problems from [Advent of Code](https://adventofcode.com)

Initial impressions are that this is quite a nice way to build up knowledge of a new language, and I might try that each year with Advent of Code.

## Running the solutions

To simply execute the code for any of the given advent days, change to the directory of the same name and execute `go run main.go`

e.g.

```bash
~/go/src/github.com/neilfenwick/advent-of-code>cd day1
~/go/src/github.com/neilfenwick/advent-of-code/day1>go run main.go
```

*_Notes for benchmarks_*

Benchmarks were just an extra - I may be curious to compare how quickly I got to optimised solutions in various languages, if I attempt the same solution in C# on .NET Core 2.1, or later (using BenchmarkDotnet?), on the same PC.

```bash
$ go version
go version go1.10.3 linux/amd64

$ cat /proc/cpuinfo | grep 'model name'
model name    : Intel(R) Core(TM) i5-6600 CPU @ 3.30GHz
```

*Beware **spoilers** at links below*

## Day 1 - Inverse Captcha

[Day 1 commentary and benchmarks](day1)

## Day 2 - Corruption Checksum

[Day 2 commentary and benchmarks](day2)

## Day 3 - Spiral Memory

[Day 3 commentary and benchmarks](day3)

## Day 4 - High-Entropy Passphrases

[Day 4 commentary and benchmarks](day4)

## Day 5 - A Maze of Twisty Trampolines, All Alike

[Day 5 commentary and benchmarks](day5)

## Day 6 - Memory Reallocation

[Day 6 commentary and benchmarks](day6)

## Day 7: Recursive Circus

[Day 7 commentary and benchmarks](day7)