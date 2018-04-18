package main

import "fmt"

func permutations(a []string, k, n int) {
	if k == n {
		fmt.Println(a)
		return
	}

	for i := k; i < n; i++ {
		// swap
		a[k], a[i] = a[i], a[k]

		permutations(a, k+1, n)

		//un-swap
		a[k], a[i] = a[i], a[k]
	}
}

func main() {
	permutations([]string{"a", "b", "c", "d"}, 0, 4)
}
