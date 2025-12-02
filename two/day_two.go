package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func expandRanges(ranges []string) ([][]int, error) {
	var result [][]int

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}

		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range: %s", r)
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid start number: %s", parts[0])
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid end number: %s", parts[1])
		}

		var group []int
		for i := start; i <= end; i++ {
			group = append(group, i)
		}

		result = append(result, group)
	}

	return result, nil
}

func readFile() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text()
	}
	log.Fatal("File is empty")
	return ""
}

func partOne(groups [][]int) []int {
	var invalidNumbers []int

	for _, group := range groups {
		for _, number := range group {
			s := strconv.Itoa(number)
			length := len(s)
			if length%2 != 0 {
				continue
			}

			half := length / 2
			if s[:half] == s[half:] {
				invalidNumbers = append(invalidNumbers, number)
			}
		}
	}

	return invalidNumbers
}

func partTwo(groups [][]int) []int {
	var invalidNumbers []int
	for _, group := range groups {
		for _, number := range group {
			s := strconv.Itoa(number)
			length := len(s)
			for patternLen := 1; patternLen <= length/2; patternLen++ {
				if length%patternLen != 0 {
					continue
				}
				pattern := s[:patternLen]
				repetitions := length / patternLen
				repeated := strings.Repeat(pattern, repetitions)
				if repeated == s {
					invalidNumbers = append(invalidNumbers, number)
					break
				}
			}
		}
	}
	return invalidNumbers
}

func main() {
	line := readFile()

	ranges := strings.Split(line, ",")

	groups, err := expandRanges(ranges)
	if err != nil {
		log.Fatalf("Error expanding ranges: %v", err)
	}

	invalidNumbersOne := partOne(groups)
	invalidNumbersTwo := partTwo(groups)

	sumOne := 0
	for _, number := range invalidNumbersOne {
		sumOne += number
	}
	fmt.Println(sumOne)

	sumTwo := 0
	for _, number := range invalidNumbersTwo {
		sumTwo += number
	}
	fmt.Println(sumTwo)
}
