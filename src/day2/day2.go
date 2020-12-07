package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData() []string {
	var entries []string

	file, err := os.Open("./src/day2/day2data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return entries
}

func isPasswordValid(letter string, min int, max int, password string) bool {
	occurrences := strings.Count(password, letter)
	return occurrences >= min && occurrences <= max
}

func isLineValid(line string) bool {
	values := strings.Split(line, " ")
	bounds := strings.Split(values[0], "-")
	min, _ := strconv.Atoi(bounds[0])
	max, _ := strconv.Atoi(bounds[1])

	return isPasswordValid(strings.TrimSuffix(values[1], ":"), min, max, values[2])
}

func breakLines(lines []string) int {
	count := 0

	for _, line := range lines {
		if isLineValid(line) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(breakLines(readData()))
}
