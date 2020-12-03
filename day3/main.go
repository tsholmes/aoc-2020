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
	lines := strings.Split(input, "\n")
	res := 0
	for i, line := range lines {
		j := (i * 3) % len(line)
		if line[j] == '#' {
			res++
		}
	}
	fmt.Println(res)
}

func part2() {
	lines := strings.Split(input, "\n")
	res := 1
	for _, s := range [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		r := 0
		for i, line := range lines {
			if i%s[1] != 0 {
				continue
			}
			j := (i / s[1] * s[0]) % len(line)
			if line[j] == '#' {
				r++
			}
		}
		res = res * r
	}
	fmt.Println(res)
}

const input = `.#..........#...#...#..#.......
.###...#.#.##..###..#...#...#..
#.....#................#...#.#.
#.....#..###.............#....#
......#.....#....#...##.....###
....#........#.#......##....#.#
..#.......##..#.#.#............
#.............#..#...#.#...#...
.#...........#.#....#..##......
......#..##..#....#....#...##..
....#.##.#####..#.##..........#
..#.#......#.#.#....#.....#....
...###.##......#..#.#...#...#..
...#..#.#..#..#.......#........
...#....#..#...........#.#.....
....#.........###.#....#...#...
....#..##.....#.##....##.#.....
........#.#.#.....#........#...
..#..#.....#.#...#.#...#.#.....
....#..........#....#....#...##
.##...#..#...##....#..#.#....#.
.#....##..#...#................
..#.###.........#.###.....#....
....#..#.......###.#...........
#...#...#.#...........#.#......
.#..#.......##.....##...#......
....####.#..#.#.#...........#..
.##...#..#..#.#....##.....#..##
...#......##....#...#.#.###....
##.#...........#.........#...#.
...........#...#...........##..
.....#....#...........#........
...#..#.........#...#....#.##..
.....##.........#...#........##
....#....#..#.#...#...##.#.....
...#.#..#...#...........#..#...
.....#.#.....#....#...#....#...
.#.............#..##..........#
..........#......#..##.....###.
..#....#........#.#.....##...#.
#..#......#.#.##......#.#.##...
.....#..#.........#...#.#.#.#.#
#.#...#.......#.#..##.##.....##
.....#......##......#.......#..
#.....#...##.#.#........#......
#..........#.#...#.......#.....
..#..#........#........#.......
...#....#....#..####.#....#...#
#.............#.....##....#..#.
##....#.....###..##....#......#
#.....#...#.#.............#....
.#.#..##..##.#..#....#.#.#...#.
.#...#..#.....#..#.#.#..#...##.
..#.#.#.#.#.#....##...#........
.......##.....#..........#...#.
...#..#...#...........#....#...
.....#..#....#..#.##...#.......
..##..#.......#.#..#....#......
...#...............#.#..#......
....#........#...#....#...#.#..
...#...#..........##....##.#...
..###.#.##.............#..#.#.#
##.......##.#..#.#.#.....#.#.#.
..#####...#......##...#........
...#.##...#................#..#
..#......#...#....#.#..##..#...
#.#.........#............#.....
##.............#.#.....#......#
....#.......#..#..##....#.#....
...#...##....#.........#..#....
...####.....#...........#....#.
#.#........##....#..#..#...#...
....#.#.###..........#........#
#.#......#.....#.##....#.#...#.
#....##.#..##..#.#.............
.#.....##..#..................#
...#.#........#...#.#........#.
..#....#......#.....##........#
....#...#....#...#.....#.##....
...#........#.......##.........
.#.##......#......#....##......
.#...#...###.#............#..#.
.#...........#.#.#....#...#..#.
.#.....#....#.....#...#........
.#..#.....#............#.#.##.#
...###.#.............#..##.....
...#.#.##.#..#..........#..#...
.#.#.#....#..#...............##
.......#.#..#...#.#.#........#.
....#.#...#..##....#........#.#
..........#...#.......#..#....#
...###.....#.#....#.....##.....
#......#..#..#........#.#...#..
#......#....#..#.#.............
...#....#........#...#..#......
...#..###........#.#.........##
#......#.#..###..#........###..
.#.#......#.#..#.#.#.#.....#..#
#....#.....#..##.....#.........
....#......#...#..#..#.#.##.#..
........#.#...#...#..#...#.#..#
.....##........#...#....#...#..
....#...##..#........#....##.#.
...............#.....#......##.
..##.....#.....#.#.............
.....#.#...........##.#.....#..
.#..##..#.##.#...##.#....#....#
.##.....#.##......#....#..#..#.
.......#.##......#....#...#.#..
.#........#......#...##.#....#.
.........#..........#.......###
#.#.........#..#..#....#...#...
.......#.........#......#.#.#..
.......#...........#....#....#.
.###...##.#.#..........#...#..#
....#.....#...#..#.............
.......##........#..#.......#..
....##..#.#....#....#..#...#..#
..#.####.....#.........#.#....#
..............#.#..#.....#...#.
.....#.............#..........#
..##.#...#.....#....#.#....##..
.#...#.......#..####..#..#...#.
#..........#................##.
......##.....#.................
..##...#.#..........##.#...#...
....#.#.#.#...##...#...#...####
.............##..#.###...#.....
#.#....#.#..#..##........#..##.
.....#.#...............#.......
...#..##......#..##...........#
#..#....#...........##..#......
.##....#.#....###.......#..#...
.....#..#.#....##...#......#...
.#.........#####......#...#...#
.......#.#.....#.....#.......#.
#....#.......###.......#..#....
#......##.###...#.......#......
.......#...#......#....#..#....
.#.####.......#...#.##.........
................##.#......#....
......##....#.#......#......#..
....##...##....#.........#.....
......#.#..............##.#...#
....#.#......#.#.............#.
.#.#..####...#................#
....#.#.#.#......##...##......#
.....#.#..#......#....#......#.
..........#.#.....#.......#...#
..##......##.#...##.#......#..#
...#............#..#...###.....
.#.#..###..#.......##...#.....#
.#....#.#.......#.....##....#..
#.............###...##.#.#...#.
#........#.#........#.#...#.#.#
##..#.................#....#...
...#.#...#..#.#..##....#...#...
#.....#.......#..............#.
.......###...##..#.....#.......
#.#.........#..#.#.........#...
.#.#............#.....##.....#.
........#....#....#.......#....
...#.#....#..#.##....#.#......#
.#.....#.#..#...........#.#.#..
#......#..#......##.#.#.#.#..#.
.......#.#..#......#.#.#..#.#.#
..........#...#..........#.##..
.#.#..####.......#..........#..
......#.#.....#..#..#..#.....#.
.....##..#.#.#..#..#...#.....##
............#.#....#.#....#....
..............#..#...#...#.....
.....#......#.......#.....#....
..##....#..#...........#..##...
###...#.##..#.#...####....###..
..#.#.....#.........#....#..###
##...........##.............#..
....##..............#.........#
...#...##....#.#..#...##.....#.
..#..##...#.......#..#..#.....#
...#...#....####........##.#...
....#........#..#.#.........#..
.#..........#...#..#.#.#......#
....#.#.....#.........#....#...
...#....#...##.......#...#.....
....#..#.......#.##.##.##...#..
##....##........#........##....
.#.#..#...........#.....#...#..
...#.##...##..#...#...##.......
.....#..###................#.#.
...#........##.#....##.....#.##
...#...#..##...#...#.#...#.....
.#......#...#..#.##.......#...#
.....#.......###.##...#........
#.....#..#........##.##.#.##..#
....#..............##.##...#...
#..........#..................#
..##.......#..........#..#..##.
.#....###.#..#.........###....#
.#....#.##..............#.##.##
.#.##.#....#.......#.#......#..
.#............#.#.....#........
..#......#.......#.............
#.#...#........##...#.#......#.
....#.........#........##..#...
..........##.....#.#......#....
.##.#..#....#.......#...#...##.
.#................#...#.##.....
....###.......#..#..#.........#
.#.....#..##...###......#.....#
.#.##..........#..#..#........#
.......#.##..............#...##
#...#.#.#.......#..#......#.##.
.#....#.#......#...#..........#
.....#........##....#.##.....#.
.#....................#..#.#.#.
.....#.........#....#.......#.#
.....#.#..##..#.....#..#.......
...#..#..#...#.....#....#....#.
#.....#.#.#..........#..#.#.#..
.....##..##.....#.#..#.........
#.#..##....##......##...#.##..#
..##..#.....#..#..........##...
......#.#...#..#.......##.....#
..#.#.......#.#......#.........
.....#........##..#.....####.#.
.#.....#........#.......#..##..
......#...#....#.##...#.......#
..##..................#..#.....
.....###.#..##...#.............
...##...##...#......#....#....#
#........#.#..........##..#....
#........#....#..........#...#.
...##.#.##..#...##......#......
#........##....#.#..##.....#..#
...####......#..#......#.#.....
.#......#...#...#.#.....##....#
.....###..##..#...#..........##
##.##....#...#.................
...##.#.......#.###......#..#..
.....#.#.#.......#.......#..#.#
#...#...#.##..#....###.......#.
.#.#..##.....#....#...##.......
.....#..........#....#...#.##..
..........#....#...#...........
.#....#..#...#...#.......#....#
#..#..............#.....####.##
.......#....###....#....#.#.#..
###.#........##.#.......#......
#..#...#..#......#.............
#...###..#...#..#..##.#.###.#..
..#..#...##......##............
.#..#.......#..###..##...#.....
....#..#..##.#.#.....##...#.#.#
....#....#.....#..#....#.......
..##..#....#.#...##..#.........
.....#....#...........#.#......
...#........#.#..#..#......#..#
.#...##....#....#.#.##......#.#
..#...........#..###.##.....#..
.#.######.#..##.......#..#.....
.....#..#......##.#.#...#......
....#....#..#.....#.......#.#.#
.....#........##.....#.....#.##
........#....#...#...#.#.#...#.
...#.#.....#...........#.....#.
#.#.#...###......#.....#.....#.
.#..........#.....#.......##...
#................#.#.....#.####
.#......#......#.#..##.#.##....
..........#....#...........###.
.##....#..####..#####..........
##.......##............#.....#.
...#.....#...#....#.......#....
.#....##......#.#...#....#.....
....#............##..........#.
.#....#....#.....#.#...........
.............##.#.##...#.#.#...
..#............#.#..##.#....##.
#.....#...##..........#.#.#...#
......#............#..........#
..##..#.....#........#.##..#..#
#..#.#..##.#.....##.#..........
#..#...#.#..#......##.......##.
.##......#...........##.....#..
...#.....#.....#..#....#.......
.....#...............#........#
.......#.....##..#..##..#.#.#..
#.#.....#..#..........##...#...
#..#......#.................#.#
.##...#....#...#...#.......#...
.#........##........#..........
........#..........#.........#.
.....#.##..#.......#........#..
..##..#..#...##..#.#....#......
......#........#.##.....#.#....
.#...#.#.........#..#.#.#.#..#.
.#..#.#...#............#.#..#..
....#.................#...#..##
.........##.....#.#.#......####
...............#....##.#.#.....
....##..#....#......#....#.....
....##.#...#....#.#..#...#..#..
..##......#.#..#........#.#.#..
.........#.#................##.
##.....#.....##..##.#........#.
###....#..#..#..#..#.##..##.#..
.....##..#...........##..#.#...
....#..#..#..#....#...#.#....#.
#....#............#..#....###..
....#..#.............#....##.#.
...#.................#...#.....
.##...#....#..#..#........#....
...#.#..#...#.#......#....#....
...#.......##..........#...#.#.
...##..#.......#........#...#..
.....#.#.#....#..##......##...#
....##......#........##....##..
..#..........#.#.##.....#......
..................#..#..#..###.
.#..............#.#..#.#..#.###
..#....#....#......#..##..#...#
#.........#..#..#...........#..`