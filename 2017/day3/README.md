# Day 3: Spiral Memory

## Part 1

### Working through the problem

This one was quite a bit trickier and went down a rabbit hole before abandoning the approach and coming back to re-assess the problem.  Lesson learned!

The problem space is a two-dimensional spiral grid and the [Manhattan Distance](https://en.wikipedia.org/wiki/Taxicab_geometry) is essentially the sum of the absolute horizontal distance and absolute vertical distance between two poIntegers.

My initial impressions of the problem were naive. I initially thought to approach it by iteratively generating a map of the poIntegers, but after trying to do that I realised the approach is a bit of a naive brute-force approach and was a bit fiddly. (I would say the brute-force solution is more error-prone, but with a TDD approach, the tests were telling me pretty much the whole time that it wasn't working.)

A more mathematical approach is to look at the spiral grid as a set of concentric squares, and if you extrapolate the grid a little more, a pattern emerges when you look at the highest values in each square-"ring" (bottom right diagonal):

```text
  37  36  35  34  33  32  31
  38  17  16  15  14  13  30
  39  18   5   4   3  12  29
  40  19   6   1   2  11  28  53
  41  20   7   8   9  10  27  52
  42  21  22  23  24  25  26  51
  43  44  45  46  47  48  49  50
                              81
```

Look in the bottom right corner of each "ring", take the square root of that number and compare it to the "width" of the "ring":

```text
  sqrt(1) = 1
  sqrt(9) = 3
  sqrt(25) = 5
  sqrt(49) = 7
```

Given an arbitrary number `n`, we can quickly work out what "ring" it is in by finding the next closest (highest) odd-number-squared.  From there it should be fairly straight forward to count the number of steps back to the centre.

```text
  (x-2)^2 < n <= (x)^2
```

Where:

- x is odd
- x greater than, or equal to, 3

To count the steps to the centre:

1. find the closest number that is on one of the axes of the cartesian plane and count the absolute value of the distance to it
2. take the square root of the largest number on the "ring", halve it and round down (this is the distance from the point where the "ring" bisects the axis to the centre of the grid)
3. add the results from the previous two steps

### Part 1 Benchmarks

```text
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day3/grid -bench ^BenchmarkManhattanDistance$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day3/grid
BenchmarkManhattanDistance-4       50000000            37.5 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day3/grid    1.916s
Success: Benchmarks passed.
```

## Part 2

I'm immediately thinking that I'm regretting not tracing the grid in a brute-force approach in the first part now because it looks like I would need to do that anyway.

I extrapolated the grid a little in an attempt to search for patterns.

```text
147  142  133  122   59
304    5    4    2   57
330   10    1    1   54
351   11   23   25   26
362  747  806  880  931
```

### Approach

I think I am going to create a stateful struct that maintains an in-memory representation of the grid in cartesian co-ordinates.  I'll then create some functions, with that struct as the receiver, that do the following:

 1. traverse the grid on a spiral path, up to a position number, returning all poIntegers on the path
 2. a function to return all the neighbours that are within a distance of sqrt(2) pythagorean distance of a point
 3. sum the cumulative values of the neighbours

### Part 2 Benchmarks

#### Pre-optimised

Wow, that timing number is large compared to some other advent solutions so far...

I think that is because the solution iterates to find the "first value written that is larger than" the input by incrementing upwards, linearly, from 1.

A binary-chop / goal-seek approach would probably be faster.

```text
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day3/grid -bench ^BenchmarkCumulativeSum$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day3/grid
BenchmarkCumulativeSum-4            500       3324585 ns/op      623047 B/op        6736 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day3/grid    2.003s
Success: Benchmarks passed.
```

#### First-optimisation - add cumulative sums to spiral state

*Much* better!  I suppose that fits with my experience that sorting and searching algorithms eventually tend towards optimising for a trade-off of memory vs cpu, at some point.

```text
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ github.com/neilfenwick/advent-of-code/day3/grid -bench ^BenchmarkCumulativeSum$

goos: linux
goarch: amd64
pkg: github.com/neilfenwick/advent-of-code/day3/grid
BenchmarkCumulativeSum-4        1000000          1790 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/neilfenwick/advent-of-code/day3/grid    1.812s
Success: Benchmarks passed.
```