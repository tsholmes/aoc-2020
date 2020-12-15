package main

import (
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	ps := strings.Split(input, ",")
	nums := make([]int, len(ps))
	for i, p := range ps {
		fmt.Sscanf(p, "%d", &nums[i])
	}
	last := map[int]int{}
	diffs := map[int]int{}
	lnum := 0
	for turn := 1; turn <= 2020; turn++ {
		if turn <= len(nums) {
			lnum = nums[turn-1]
			v, ok := last[lnum]
			if ok {
				diffs[lnum] = turn - v
			}
			last[lnum] = turn
		} else {
			v, ok := diffs[lnum]
			if ok {
				lnum = v
			} else {
				lnum = 0
			}
			v, ok = last[lnum]
			if ok {
				diffs[lnum] = turn - v
			}
			last[lnum] = turn
		}
	}
	fmt.Println(lnum)
}

func part2() {
	ps := strings.Split(input, ",")
	nums := make([]int, len(ps))
	for i, p := range ps {
		fmt.Sscanf(p, "%d", &nums[i])
	}
	last := map[int]int{}
	diffs := map[int]int{}
	lnum := 0
	for turn := 1; turn <= 30000000; turn++ {
		if turn <= len(nums) {
			lnum = nums[turn-1]
			v, ok := last[lnum]
			if ok {
				diffs[lnum] = turn - v
			}
			last[lnum] = turn
		} else {
			v, ok := diffs[lnum]
			if ok {
				lnum = v
			} else {
				lnum = 0
			}
			v, ok = last[lnum]
			if ok {
				diffs[lnum] = turn - v
			}
			last[lnum] = turn
		}
	}
	fmt.Println(lnum)
}

const input = `19,20,14,0,9,1`
