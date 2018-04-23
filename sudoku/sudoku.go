package main

import (
	"fmt"
	"math"
)

func getMapSet() map[int]bool {
	m := make(map[int]bool)
	for i := 1; i <= 9; i++ {
		m[i] = true
	}

	return m
}

func toSlice(m map[int]bool) []int {
	var vs []int
	for k, v := range m {
		if !v {
			continue
		}

		vs = append(vs, k)
	}

	return vs
}

// line returns the possible values for a given point by checking the column/row
func line(b [][]int, i int, row bool) (pv []int) {
	for k := 0; k < 9; k++ {
		var v int
		if row {
			v = b[i][k]
		} else {
			v = b[k][i]
		}
		if v == 0 {
			continue
		}

		pv = append(pv, v)
	}

	return pv
}

func sectPoint(i, j int) (r, c int) {
	return 3 * (i / 3), 3 * (j / 3)
}

func sector(b [][]int, x, y int) (pv []int) {
	x, y = sectPoint(x, y)
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			v := b[i][j]
			if v == 0 {
				continue
			}

			pv = append(pv, v)
		}
	}

	return pv
}

func pvs(b [][]int, i, j int) []int {
	var pv []int
	pv = append(pv, line(b, i, true)...)
	pv = append(pv, line(b, j, false)...)
	pv = append(pv, sector(b, i, j)...)

	m := getMapSet()
	for _, v := range pv {
		m[v] = false
	}

	return toSlice(m)
}

func findFree(b [][]int) (x, y int, pv []int, err error) {
	l := math.MaxInt32
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] != 0 {
				continue
			}

			gv := pvs(b, i, j)
			if len(gv) < 1 {
				return i, j, gv, fmt.Errorf("no possible values, backtrack now")
			}

			if len(gv) >= l {
				continue
			}

			l, x, y, pv = len(gv), i, j, gv
		}
	}

	return x, y, pv, nil
}

func printSolution(b [][]int) {
	for _, v := range b {
		fmt.Println(v)
	}
}

var finish = false

func solve(b [][]int, fv int) {
	if fv == 0 {
		printSolution(b)
		finish = true
		return
	}

	i, j, pv, err := findFree(b)
	if err != nil {
		return
	}

	for _, v := range pv {
		b[i][j] = v
		solve(b, fv-1)
		b[i][j] = 0
		if finish {
			return
		}
	}
}

func main() {
	// https://www.telegraph.co.uk/news/science/science-news/9359579/Worlds-hardest-sudoku-can-you-crack-it.html
	b := [][]int{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}

	var n int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := b[i][j]
			if v == 0 {
				n++
			}
		}
	}

	solve(b, n)
}
