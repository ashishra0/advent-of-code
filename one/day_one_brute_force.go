package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var (
// 	currentPosition    int  = 50
// 	first_time_attempt bool = true
// 	count              int  = 0
// )

// func readInput() []string {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatalf("Error opening file: %v", err)
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatalf("Error reading file: %v", err)
// 	}
// 	return lines
// }

// func execute(line string) {
// 	if strings.HasPrefix(line, "R") {
// 		addition(line)
// 	} else if strings.HasPrefix(line, "L") {
// 		subtraction(line)
// 	}
// }

// func addition(line string) {
// 	number, err := strconv.Atoi(line[1:])
// 	if err != nil {
// 		log.Fatalf("Error converting string to int: %v", err)
// 	}

// 	currentPosition = currentPosition + number
// 	currentPosition = currentPosition % 100

// 	if first_time_attempt {
// 		first_time_attempt = false
// 	}

// 	if currentPosition == 0 {
// 		count++
// 		fmt.Printf("HIT ZERO at count=%d, line=%s, from addition\n", count, line)
// 	}
// }

// func subtraction(line string) {
// 	number, err := strconv.Atoi(line[1:])
// 	if err != nil {
// 		log.Fatalf("Error converting string to int: %v", err)
// 	}

// 	if first_time_attempt {
// 		currentPosition = 100 - (number - currentPosition)
// 		first_time_attempt = false
// 	} else {
// 		if number > currentPosition {
// 			currentPosition = 100 - (number - currentPosition)
// 		} else if number < currentPosition {
// 			currentPosition = currentPosition - number
// 		} else {
// 			currentPosition = 0
// 		}
// 	}

// 	// Always ensure position stays in 0-99 range
// 	currentPosition = ((currentPosition % 100) + 100) % 100

// 	if currentPosition == 0 {
// 		count++
// 	}
// }

// func main() {
// 	lines := readInput()
// 	for _, line := range lines {
// 		execute(line)
// 	}
// 	fmt.Printf("Count: %d\n", count)
// }
