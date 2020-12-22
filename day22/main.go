package main

import (
	"fmt"

	"github.com/tsholmes/aoc-2020"
)

func main() {
	part1()
	part2()
}

func part1() {
	secs := aoc.LineSections(input)

	p1 := make([]int, len(secs[0])-1)
	for i, line := range secs[0][1:] {
		fmt.Sscanf(line, "%d", &p1[i])
	}
	p2 := make([]int, len(secs[1])-1)
	for i, line := range secs[1][1:] {
		fmt.Sscanf(line, "%d", &p2[i])
	}

	for len(p1) > 0 && len(p2) > 0 {
		var c1, c2 int
		c1, p1 = p1[0], p1[1:]
		c2, p2 = p2[0], p2[1:]
		if c1 > c2 {
			p1 = append(p1, c1, c2)
		} else {
			p2 = append(p2, c2, c1)
		}
	}

	res := 0
	wd := p1
	if len(p1) == 0 {
		wd = p2
	}
	for i, c := range wd {
		res += c * (len(wd) - i)
	}

	fmt.Println(res)
}

func part2() {
	secs := aoc.LineSections(input)

	p1 := make([]byte, len(secs[0])-1)
	for i, line := range secs[0][1:] {
		fmt.Sscanf(line, "%d", &p1[i])
	}
	p2 := make([]byte, len(secs[1])-1)
	for i, line := range secs[1][1:] {
		fmt.Sscanf(line, "%d", &p2[i])
	}

	var play func([]byte, []byte) (bool, []byte, []byte)
	play = func(p1 []byte, p2 []byte) (bool, []byte, []byte) {
		seen := map[string]bool{}
		p1, p2 = append([]byte{}, p1...), append([]byte{}, p2...)

		for len(p1) > 0 && len(p2) > 0 {
			k := string(p1) + "z" + string(p2)
			if seen[k] {
				return false, p1, p2
			}
			seen[k] = true

			var c1, c2 byte
			c1, p1 = p1[0], p1[1:]
			c2, p2 = p2[0], p2[1:]

			rwinner := false
			if len(p1) < int(c1) || len(p2) < int(c2) {
				rwinner = c2 > c1
			} else {
				rwinner, _, _ = play(p1[:c1], p2[:c2])
			}
			if rwinner {
				p2 = append(p2, c2, c1)
			} else {
				p1 = append(p1, c1, c2)
			}
		}

		return len(p2) != 0, p1, p2
	}

	winner, p1, p2 := play(p1, p2)
	res := 0
	wd := p1
	if winner {
		wd = p2
	}
	for i, c := range wd {
		res += int(c) * (len(wd) - i)
	}

	fmt.Println(res)
}

const input = `Player 1:
1
10
28
29
13
11
35
7
43
8
30
25
4
5
17
32
22
39
50
46
16
26
45
38
21

Player 2:
19
40
2
12
49
23
34
47
9
14
20
24
42
37
48
44
27
6
33
18
15
3
36
41
31`
