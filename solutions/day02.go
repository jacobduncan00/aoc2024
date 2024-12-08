package solutions

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/jacobduncan00/aoc2024/utils"
)

func Day2(input string) (interface{}, interface{}) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	safeCount := 0
	dampenerSafeCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			fmt.Println("Line is empty")
			continue
		}

		var levels []int
		for _, levelStr := range strings.Fields(line) {
			levelNum, err := strconv.Atoi(levelStr)
			if err != nil {
				fmt.Println("Error converting")
				continue
			}
			levels = append(levels, levelNum)
		}

		if isReportSafe(levels) {
			safeCount++
			dampenerSafeCount++
			continue
		}

		if isReportSafeWithDampener(levels) {
			dampenerSafeCount++
		}

	}

	return safeCount, dampenerSafeCount
}

func isReportSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	firstDiff := levels[1] - levels[0]
	if firstDiff == 0 {
		return false
	}

	isIncreasing := firstDiff > 0

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		absDiff := utils.Abs(diff)
		if absDiff < 1 || absDiff > 3 {
			return false
		}

		if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) {
			return false
		}
	}

	return true
}

func isReportSafeWithDampener(levels []int) bool {
	for i := range levels {
		dampened := make([]int, 0, len(levels)-1)
		dampened = append(dampened, levels[:i]...)
		dampened = append(dampened, levels[i+1:]...)

		if isReportSafe(dampened) {
			return true
		}
	}

	return false
}
