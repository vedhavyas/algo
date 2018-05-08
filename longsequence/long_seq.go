package main

import "fmt"

func findLongestSeq(a []int) (count int, path []int) {
	if len(a) == 0 {
		return count, path

	}

	path = append(path, a[0])
	count++

	for i := 1; i < len(a); i++ {
		switch {
		case a[i] > path[len(path)-1]:
			path = append(path, a[i])
			count++
		case a[i] < path[len(path)-1]:
			if len(path) < 2 || a[i] > path[len(path)-2] {
				path[len(path)-1] = a[i]
			}
		}
	}

	return count, path
}

func main() {
	fmt.Println(findLongestSeq([]int{2, 4, 3, 5, 1, 7, 6, 9, 8}))
}
