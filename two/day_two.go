package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {
	lines := readFile()
	fmt.Println(lines)
}
