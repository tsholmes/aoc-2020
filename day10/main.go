package main

import (
	"fmt"
	"sort"

	"github.com/tsholmes/aoc-2020"
)

func main() {
	part1()
	part2()
}

func part1() {
	nums := aoc.NumLines(input)
	sort.Ints(nums)
	last := 0
	c1, c3 := 0, 1
	for _, n := range nums {
		switch n - last {
		case 1:
			c1++
		case 3:
			c3++
		}
		last = n
	}
	fmt.Println(c1 * c3)
}

func part2() {
	nums := aoc.NumLines(input)
	sort.Ints(nums)
	counts := map[int]int{}
	counts[0] = 1
	for _, n := range nums {
		for i := -3; i <= -1; i++ {
			counts[n] += counts[n+i]
		}
	}
	fmt.Println(counts[nums[len(nums)-1]])
}

const input = `104
83
142
123
87
48
102
159
122
69
127
151
147
64
152
90
117
132
63
109
27
47
7
52
59
11
161
12
148
155
129
10
135
17
153
96
3
93
82
55
34
65
89
126
19
72
20
38
103
146
14
105
53
77
120
39
46
24
139
95
140
33
21
84
56
1
32
31
28
4
73
128
49
18
62
81
66
121
54
160
158
138
94
43
2
114
111
110
78
13
99
108
141
40
25
154
26
35
88
76
145`
