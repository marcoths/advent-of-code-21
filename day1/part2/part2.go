package part2

func Part2Answer(input []int) int {

	var measurements []int
	windowSum, windowStart := 0, 0

	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		windowSum += input[windowEnd]
		if windowEnd >= 2 {
			measurements = append(measurements, windowSum)
			windowSum -= input[windowStart]
			windowStart++
		}
	}

	var answer int
	for i, v := range measurements {
		if i > 0 && v > measurements[i-1] {
			answer += 1
		}
	}

	return answer

}
