package main

import "fmt"

func bc(a [][]int, m, n int) {
	if m == n || n == 0 {
		a[m][n] = 1
		return
	}

	a[m][n] = a[m-1][n-1] + a[m-1][n]
}

func binomialCoefficientOf(m, n int) int {
	m, n = m+1, n+1
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i; j < m; j++ {
			bc(a, j, i)
		}
	}

	return a[m-1][n-1]
}

func main() {
	fmt.Println(binomialCoefficientOf(10, 5))
}
