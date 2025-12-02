package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return lines
}

func partOne(lines []string) {
	position := 50
	count := 0

	for _, line := range lines {
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		if strings.HasPrefix(line, "L") {
			position = position - number
		} else {
			position = position + number
		}

		position = ((position % 100) + 100) % 100

		if position == 0 {
			count++
		}
	}

	fmt.Printf("Count: %d\n", count)
}

func partTwo(lines []string) {
	position := 50
	count := 0

	for _, line := range lines {
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		for i := 0; i < number; i++ {
			if strings.HasPrefix(line, "L") {
				position = position - 1
			} else {
				position = position + 1
			}

			position = ((position % 100) + 100) % 100

			if position == 0 {
				count++
			}
		}
	}

	fmt.Printf("Count: %d\n", count)
}

func main() {
	lines := readFile()
	partOne(lines)
	partTwo(lines)
}
