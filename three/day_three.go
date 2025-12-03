package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		var numbersInt []int
		for _, char := range line {
			num := int(char - '0')
			numbersInt = append(numbersInt, num)
		}

		result = append(result, numbersInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return result
}

func partOne(lines [][]int) []int {
	var maxPairs []int

	for _, line := range lines {
		if len(line) < 2 {
			continue
		}

		maxPair := 0

		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				pair := line[i]*10 + line[j]
				if pair > maxPair {
					maxPair = pair
				}
			}
		}

		maxPairs = append(maxPairs, maxPair)
	}

	return maxPairs
}

func partTwo(lines [][]int) []int64 {
	var maxNumbers []int64

	for _, line := range lines {
		if len(line) < 12 {
			continue
		}

		var chosen []int
		pos := 0

		for len(chosen) < 12 {
			digitsStillNeeded := 12 - len(chosen)
			digitsAvailable := len(line) - pos

			canSkip := digitsAvailable - digitsStillNeeded

			maxDigit := line[pos]
			maxPos := pos

			for i := pos; i <= pos+canSkip; i++ {
				if line[i] > maxDigit {
					maxDigit = line[i]
					maxPos = i
				}
			}

			chosen = append(chosen, maxDigit)
			pos = maxPos + 1
		}

		var number int64 = 0
		for _, digit := range chosen {
			number = number*10 + int64(digit)
		}

		maxNumbers = append(maxNumbers, number)
	}

	return maxNumbers
}

func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	lines := readFile()
	maxPairs := partOne(lines)
	fmt.Println(sum(maxPairs))
	maxNumbers := partTwo(lines)
	sum := 0
	for _, number := range maxNumbers {
		sum += int(number)
	}
	fmt.Println(sum)
}
