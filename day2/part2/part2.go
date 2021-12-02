package part2

import (
	"strconv"
	"strings"
)

func Part2Answer(input string) int {
	inputList := strings.Split(input, "\n")
	var position int
	var depth int
	var aim int

	for _, v := range inputList {
		line := strings.Split(v, " ")
		direction := line[0]
		value, _ := strconv.Atoi(line[1])

		switch direction {
		case "forward":
			position += value
			depth += (aim * value)
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

	return position * depth
}
