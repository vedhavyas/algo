package main

import "fmt"

const (
	Match = iota
	Insert
	Delete
)

type cell struct {
	cost   int
	parent int
}

func makeMatrix(i, j int) [][]cell {
	m := make([][]cell, i+1)
	for r := 0; r <= i; r++ {
		m[r] = make([]cell, j+1)
	}

	return m
}

func rowInt(m [][]cell, c int) {
	for j := 0; j <= c; j++ {
		m[0][j].cost = j
		m[0][j].parent = -1
		if j > 0 {
			m[0][j].parent = Insert
		}

	}
}

func columnInt(m [][]cell, c int) {
	for j := 0; j <= c; j++ {
		m[j][0].cost = j
		m[j][0].parent = -1
		if j > 0 {
			m[j][0].parent = Delete
		}
	}
}

func match(p, t uint8) int {
	if p == t {
		return 0
	}

	return 1
}

func inDel(c uint8) int {
	return 1
}

func minDistanceEdit(p, t string) (c int, m [][]cell) {
	m = makeMatrix(len(p), len(t))
	rowInt(m, len(t))
	columnInt(m, len(p))

	for i := 1; i <= len(p); i++ {
		for j := 1; j <= len(t); j++ {
			var a [3]int
			a[Match] = m[i-1][j-1].cost + match(p[i-1], t[j-1])
			a[Insert] = m[i][j-1].cost + inDel(t[j-1])
			a[Delete] = m[i-1][j].cost + inDel(p[i-1])

			c := a[Match]
			ac := Match
			for p := Insert; p <= Delete; p++ {
				if a[p] < c {
					c = a[p]
					ac = p
				}
			}

			m[i][j].cost = c
			m[i][j].parent = ac
		}
	}

	return m[len(p)][len(t)].cost, m
}

func reconstructPath(m [][]cell, p, t string, i, j int) {
	a := m[i][j].parent
	switch a {
	case Match:
		reconstructPath(m, p, t, i-1, j-1)
		if p[i-1] == t[j-1] {
			fmt.Print("M")
		} else {
			fmt.Print("S")
		}
	case Insert:
		reconstructPath(m, p, t, i, j-1)
		fmt.Print("I")
	case Delete:
		reconstructPath(m, p, t, i-1, j)
		fmt.Print("D")
	}
}

func main() {
	p := "thou"
	t := "you"
	c, m := minDistanceEdit(p, t)
	fmt.Println("Cost:", c)
	for i := 0; i < len(p)+1; i++ {
		for j := 0; j < len(t)+1; j++ {
			fmt.Print(m[i][j].cost, " ")
		}
		fmt.Println()
	}

	fmt.Println("Parent Matrix:")
	for i := 0; i < len(p)+1; i++ {
		for j := 0; j < len(t)+1; j++ {
			fmt.Print(m[i][j].parent, " ")
		}
		fmt.Println()
	}
	reconstructPath(m, p, t, len(p), len(t))
}
