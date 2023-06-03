package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func insertAndFlushFrom(index int, elves []int, calories int) int {
	elves[index] = calories
	return index + 1
}

func main() {
	file, err := os.Open("adventofcode.com_2022_day_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currElf := 0
	elves := make([]int, 0)
	for scanner.Scan() {
		calories := scanner.Text()

		if calories == "" {
			elves = append(elves, currElf)
			currElf = 0
			continue
		}

		caloriesInt, err := strconv.Atoi(calories)
		if err != nil {
			log.Fatal(err)
		}

		currElf += caloriesInt
	}

	if currElf > 0 {
		elves = append(elves, currElf)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(elves)

	len := len(elves) - 1

	fmt.Println("Top G", elves[len])
	fmt.Println("Top 3 Gs combined", elves[len]+elves[len-1]+elves[len-2])
}
