package solutions

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"github.com/jacobduncan00/aoc2024/utils"
)

func Day1(input string) (interface{}, interface{}) {
	var leftList, rightList []int
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}

		left, err1 := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sortedLeft := make([]int, len(leftList))
	sortedRight := make([]int, len(rightList))
	copy(sortedLeft, leftList)
	copy(sortedRight, rightList)

	sort.Ints(sortedLeft)
	sort.Ints(sortedRight)

	totalDistance := 0
	for i := 0; i < len(sortedLeft) && i < len(sortedRight); i++ {
		distance := utils.Abs(sortedLeft[i] - sortedRight[i])
		totalDistance += distance
	}

	rightNumCounts := make(map[int]int)
	for _, num := range rightList {
		rightNumCounts[num]++
	}

	similarityScore := 0
	for _, leftNum := range leftList {
		count := rightNumCounts[leftNum]
		similarityScore += leftNum * count
	}

	return totalDistance, similarityScore
}
