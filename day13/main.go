package main

import (
	"fmt"
	"strings"

	"github.com/tsholmes/aoc-2020"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := aoc.Lines(input)
	var t int
	fmt.Sscanf(lines[0], "%d", &t)
	ps := strings.Split(lines[1], ",")
	min := t * 10
	mid := 0
	for _, p := range ps {
		if p == "x" {
			continue
		}
		var x int
		fmt.Sscanf(p, "%d", &x)
		tm := t % x
		if tm == 0 {
			tm = t
		} else {
			tm = t + (x - tm)
		}
		if tm < min {
			min = tm
			mid = x
		}
	}
	fmt.Println((min - t) * mid)
}

func part2() {
	lines := aoc.Lines(input)
	ps := strings.Split(lines[1], ",")
	var pis [][2]int
	for i, p := range ps {
		if p == "x" {
			continue
		}
		var x int
		fmt.Sscanf(p, "%d", &x)
		pis = append(pis, [2]int{i, x})
	}

	pos := 0
	offset := 1
	for _, p := range pis {
		i, id := p[0], p[1]
		for ((pos + i) % id) != 0 {
			pos += offset
		}
		offset *= id
	}
	fmt.Println(pos)
}

const input = `1003240
19,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,787,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,x,x,x,x,x,x,23,x,x,x,x,x,29,x,571,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17`
