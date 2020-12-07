// parsing functions for speed
package aoc

import (
	"fmt"
	"strings"
)

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func NumLines(input string) []int {
	lines := strings.Split(input, "\n")
	nums := make([]int, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "%d", &nums[i])
	}
	return nums
}

func LineSections(input string) [][]string {
	rawsecs := strings.Split(input, "\n\n")
	secs := make([][]string, len(rawsecs))
	for i, sec := range rawsecs {
		secs[i] = strings.Split(sec, "\n")
	}
	return secs
}

func FieldSections(input string) [][]string {
	rawsecs := strings.Split(input, "\n\n")
	secs := make([][]string, len(rawsecs))
	for i, sec := range rawsecs {
		secs[i] = strings.Fields(sec)
	}
	return secs
}

func NumSections(input string) [][]int {
	rawsecs := strings.Split(input, "\n\n")
	secs := make([][]int, len(rawsecs))
	for i, sec := range rawsecs {
		ps := strings.Fields(sec)
		secs[i] = make([]int, len(ps))
		for j, p := range ps {
			fmt.Sscanf(p, "%d", &secs[i][j])
		}
	}
	return secs
}
