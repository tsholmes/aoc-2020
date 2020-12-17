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
	lines := aoc.Lines(input)
	rawlines := [][][]byte{make([][]byte, len(lines))}
	for i, line := range lines {
		rawlines[0][i] = []byte(line)
	}
	get := func(state [][][]byte, x, y, z int) bool {
		xx, yy, zz := len(state), len(state[0]), len(state[0][0])
		if x < 0 || x >= xx || y < 0 || y >= yy || z < 0 || z >= zz {
			return false
		}
		return state[x][y][z] == '#'
	}
	sim := func(state [][][]byte) [][][]byte {
		xx, yy, zz := len(state), len(state[0]), len(state[0][0])
		next := make([][][]byte, xx+2)
		for x := 0; x < xx+2; x++ {
			next[x] = make([][]byte, yy+2)
			for y := 0; y < yy+2; y++ {
				next[x][y] = make([]byte, zz+2)
				for z := 0; z < zz+2; z++ {
					count := 0
					for xo := -1; xo <= 1; xo++ {
						for yo := -1; yo <= 1; yo++ {
							for zo := -1; zo <= 1; zo++ {
								if xo == 0 && yo == 0 && zo == 0 {
									continue
								}
								if get(state, x-1+xo, y-1+yo, z-1+zo) {
									count++
								}
							}
						}
					}
					if get(state, x-1, y-1, z-1) {
						if count == 2 || count == 3 {
							next[x][y][z] = '#'
						} else {
							next[x][y][z] = '.'
						}
					} else {
						if count == 3 {
							next[x][y][z] = '#'
						} else {
							next[x][y][z] = '.'
						}
					}
				}
			}
		}
		return next
	}
	for i := 0; i < 6; i++ {
		rawlines = sim(rawlines)
	}
	res := 0
	for _, r := range rawlines {
		for _, r2 := range r {
			for _, r3 := range r2 {
				if r3 == '#' {
					res++
				}
			}
		}
	}
	fmt.Println(res)
}

func part2() {
	lines := aoc.Lines(input)
	rawlines := [][][][]byte{{make([][]byte, len(lines))}}
	for i, line := range lines {
		rawlines[0][0][i] = []byte(line)
	}
	get := func(state [][][][]byte, x, y, z, w int) bool {
		xx, yy, zz, ww := len(state), len(state[0]), len(state[0][0]), len(state[0][0][0])
		if x < 0 || x >= xx || y < 0 || y >= yy || z < 0 || z >= zz || w < 0 || w >= ww {
			return false
		}
		return state[x][y][z][w] == '#'
	}
	sim := func(state [][][][]byte) [][][][]byte {
		xx, yy, zz, ww := len(state), len(state[0]), len(state[0][0]), len(state[0][0][0])
		next := make([][][][]byte, xx+2)
		for x := 0; x < xx+2; x++ {
			next[x] = make([][][]byte, yy+2)
			for y := 0; y < yy+2; y++ {
				next[x][y] = make([][]byte, zz+2)
				for z := 0; z < zz+2; z++ {
					next[x][y][z] = make([]byte, ww+2)
					for w := 0; w < ww+2; w++ {
						count := 0
						for xo := -1; xo <= 1; xo++ {
							for yo := -1; yo <= 1; yo++ {
								for zo := -1; zo <= 1; zo++ {
									for wo := -1; wo <= 1; wo++ {
										if xo == 0 && yo == 0 && zo == 0 && wo == 0 {
											continue
										}
										if get(state, x-1+xo, y-1+yo, z-1+zo, w-1+wo) {
											count++
										}
									}
								}
							}
						}
						if get(state, x-1, y-1, z-1, w-1) {
							if count == 2 || count == 3 {
								next[x][y][z][w] = '#'
							} else {
								next[x][y][z][w] = '.'
							}
						} else {
							if count == 3 {
								next[x][y][z][w] = '#'
							} else {
								next[x][y][z][w] = '.'
							}
						}
					}
				}
			}
		}
		return next
	}
	for i := 0; i < 6; i++ {
		rawlines = sim(rawlines)
	}
	res := 0
	for _, r := range rawlines {
		for _, r2 := range r {
			for _, r3 := range r2 {
				for _, r4 := range r3 {
					if r4 == '#' {
						res++
					}
				}
			}
		}
	}
	fmt.Println(res)
}

const input2 = `.#.
..#
###`

const input = `.###.#.#
####.#.#
#.....#.
####....
#...##.#
########
..#####.
######.#`
