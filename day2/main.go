package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PlayerChoice string
type RoundOutcome string

const (
	ROCK     PlayerChoice = "A"
	PAPER    PlayerChoice = "B"
	SCISSORS PlayerChoice = "C"
)

const (
	LOSE RoundOutcome = "X"
	DRAW RoundOutcome = "Y"
	WIN  RoundOutcome = "Z"
)

func determineScore(opponent PlayerChoice, outcome RoundOutcome) int {
	score := 0

	switch outcome {
	case LOSE:
		if opponent == ROCK {
			score += 3
		} else if opponent == PAPER {
			score += 1
		} else if opponent == SCISSORS {
			score += 2
		}
	case DRAW:
		if opponent == ROCK {
			score += 1
		} else if opponent == PAPER {
			score += 2
		} else if opponent == SCISSORS {
			score += 3
		}
		score += 3
	case WIN:
		if opponent == ROCK {
			score += 2
		} else if opponent == PAPER {
			score += 3
		} else if opponent == SCISSORS {
			score += 1
		}
		score += 6
	}

	return score
}

func main() {
	file, err := os.Open("adventofcode.com_2022_day_2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()

		choices := strings.Split(line, " ")

		totalScore += determineScore(PlayerChoice(choices[0]), RoundOutcome(choices[1]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total score:", totalScore)
}
