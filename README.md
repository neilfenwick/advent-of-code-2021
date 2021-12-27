# Advent of code 2021

Brushing up my Go code skills. I've started reading "Effective Go" and probably write more idiomatic Go code than I did a year ago. This
year has generally felt easier to write in Go.

⚠ Not everything is perfectly tested.  I've used tests where they helped me think my way through a problem.  I have also intentionally ignored some error conditions because I know that the advent-of-code input is clean. I know that I can get away with that here, whereas I would have needed to be more defensive if the input was not so controlled.

Otherwise, this year, I'm really just trying to solve the Advent-of-code problems and have a little fun learning while I do it.

Interesting new things that I've learned:

- [Why pointers to slices can be useful][slicepointers]
- [Waitgroups][waitgroups] are an easier way to keep track of when concurrent goroutines complete
- How to write more readable Go - [Workshop: Practical Go - GoSG Meetup - Dave Cheney][workshop]

[slicepointers]: https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b
[waitgroups]: https://gobyexample.com/waitgroups
[workshop]: https://www.youtube.com/watch?v=gi7t6Pl9rxE