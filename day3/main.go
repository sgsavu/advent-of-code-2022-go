package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getScore(input byte) int {
	if input < 91 {
		return int(input) - 38
	}

	return int(input) - 96
}

func part1() {
	file, err := os.Open("adventofcode.com_2022_day_3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		letterCountMap := make(map[byte]int)

		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		for i := 0; i < len(firstHalf); i++ {
			letterCountMap[firstHalf[i]]++
		}

		var sus byte
		for i := 0; i < len(secondHalf); i++ {
			_, ok := letterCountMap[secondHalf[i]]
			if ok {
				sus = secondHalf[i]
			}
		}

		sum += getScore(sus)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func part2() {
	file, err := os.Open("adventofcode.com_2022_day_3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const elfCountMax = 2
	const elfCountMin = 0

	elfCount := elfCountMin

	arr := make([]map[byte]int, elfCountMax)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		letterCountMap := make(map[byte]int)
		for i := 0; i < len(line); i++ {

			if elfCount > 0 {
				prevElf := arr[elfCount-1]
				_, ok := prevElf[line[i]]
				if ok {
					letterCountMap[line[i]]++
				}
			} else {
				letterCountMap[line[i]]++
			}
		}

		if elfCount == elfCountMax {
			for i, _ := range letterCountMap {
				sum += getScore(i)
			}
			elfCount = elfCountMin
			continue
		}

		arr[elfCount] = letterCountMap
		elfCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func main() {
	part1()
	// part2()
}
