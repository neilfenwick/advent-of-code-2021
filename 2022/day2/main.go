package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		file *os.File
		err  error
	)

	strategy = scoreAsMove

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
		scoringStrategy, _ := strconv.Atoi(os.Args[2])
		if scoringStrategy > 0 {
			strategy = scoreAsResult
		}
		fallthrough
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	rounds := readRounds(file)
	var result int

	for _, rnd := range rounds {
		result += rnd.score
	}
	fmt.Println(result)
}

var strategy scoringStrategy

type scoringStrategy func(round game) int

const (
	Rock = iota
	Paper
	Scissors
	Lose
	Draw
	Win
)

func readRounds(file io.Reader) []game {
	games := []game{}
	s := bufio.NewScanner(file)
	for s.Scan() {
		plays := s.Text()
		currentGame := parsePlays(plays)
		games = append(games, currentGame)
	}
	return games
}

func parsePlays(rnd string) game {
	choices := strings.Split(rnd, " ")
	oppMove := parseMove(choices[0])
	myMove := parseMove(choices[1])
	round := game{opponent: oppMove, mine: myMove}
	round.score = strategy(round)
	return round
}

func parseMove(input string) move {
	var result move
	switch input {
	case "A", "X":
		result = Rock
	case "B", "Y":
		result = Paper
	case "C", "Z":
		result = Scissors
	}
	return result
}

func scoreAsMove(round game) int {
	var result int
	switch round.mine {
	case Rock:
		switch round.opponent {
		case Rock:
			result = 1 + 3
		case Paper:
			result = 1 + 0
		case Scissors:
			result = 1 + 6
		}
	case Paper:
		switch round.opponent {
		case Rock:
			result = 2 + 6
		case Paper:
			result = 2 + 3
		case Scissors:
			result = 2 + 0
		}
	case Scissors:
		switch round.opponent {
		case Rock:
			result = 3 + 0
		case Paper:
			result = 3 + 6
		case Scissors:
			result = 3 + 3
		}
	}
	return result
}

func scoreAsResult(round game) int {
	var result int
	switch round.opponent {
	case Rock:
		switch round.mine + 3 {
		case Draw:
			result = 1 + 3
		case Lose:
			result = 3 + 0
		case Win:
			result = 2 + 6
		}
	case Paper:
		switch round.mine + 3 {
		case Win:
			result = 3 + 6
		case Draw:
			result = 2 + 3
		case Lose:
			result = 1 + 0
		}
	case Scissors:
		switch round.mine + 3 {
		case Lose:
			result = 2 + 0
		case Win:
			result = 1 + 6
		case Draw:
			result = 3 + 3
		}
	}
	return result
}

type move int

type game struct {
	opponent move
	mine     move
	score    int
}
