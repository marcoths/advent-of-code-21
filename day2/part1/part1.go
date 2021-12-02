package part1

import (
	"strconv"
	"strings"
)

func Part1Answer(input string) int {
	inputList := strings.Split(input, "\n")
	var position int
	var depth int

	for _, v := range inputList {
		line := strings.Split(v, " ")
		direction := line[0]
		value, _ := strconv.Atoi(line[1])

		switch direction {
		case "forward":
			position += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}

	return position * depth
}
