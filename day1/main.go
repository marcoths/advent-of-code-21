package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/marcoths/advent-of-code-21/day1/part1"
	"github.com/marcoths/advent-of-code-21/day1/part2"
)

func main() {
	input, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1 answer: %d\n", part1.Part1Answer(input))
	fmt.Printf("Part 2 answer: %d\n", part2.Part2Answer(input))
}
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, value)
	}
	return lines, scanner.Err()
}
