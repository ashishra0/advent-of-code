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

func main() {
	line := readFile()

	ranges := strings.Split(line, ",")

	groups, err := expandRanges(ranges)
	if err != nil {
		log.Fatalf("Error expanding ranges: %v", err)
	}

	fmt.Println(groups)
}
