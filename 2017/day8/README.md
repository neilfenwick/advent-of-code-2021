# Day 8: I heard you like registers?

[Day 8 of advent of code 2017][2]

I ended up googling [the difference between parsers and lexers][1] when I began to write this description.  If the grammar of the input was more complicated than the example input instructions given:

```bash
b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
```

... then building a parser/lexer is probably the right approach.  However here we can just get away with using the `bufio.scanner` as a simple parser to tokenize the input lines for us.

```
These instructions would be processed as follows:

Because a starts at 0, it is not greater than 1, and so b is not modified.
a is increased by 1 (to 1) because b is less than 5 (it is 0).
c is decreased by -10 (to 10) because a is now greater than or equal to 1 (it is 1).
c is increased by -20 (to -10) because c is equal to 10.
```

[1]: https://stackoverflow.com/questions/2842809/lexers-vs-parsers
[2]: https://adventofcode.com/2017/day/8
[3]: https://github.com/campoy/advent-of-code-2018/blob/master/day01-p1/main.go