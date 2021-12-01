package part1

func Part1Answer(input []int) int {
	var answer int
	for i, v := range input {
		if i > 0 && v > input[i-1] {
			answer += 1
		}
	}
	return answer
}
