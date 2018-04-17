package main

import "fmt"

func iterate(a []string) {
	l := len(a)
	pow := 1 << uint(l)
	for i := 0; i < pow; i++ {
		fmt.Print("[")
		for j := 0; j < l; j++ {
			if i|1<<uint(j) == i {
				fmt.Print(a[j])
			}
		}
		fmt.Println("]")
	}
}

func backtrack(a []string, s []bool, k, n int) {
	if k == n {
		fmt.Print("[")
		for i, c := range s {
			if !c {
				continue
			}

			fmt.Print(a[i])
		}
		fmt.Println("]")
		return
	}

	c := []bool{true, false}
	i := k
	k += 1
	for _, v := range c {
		s[i] = v
		backtrack(a, s, k, n)
	}
}

func main() {
	fmt.Println("Iterative subset generation:")
	iterate([]string{"a", "b", "c", "d"})

	fmt.Println("\nBacktrack subset generation")
	backtrack([]string{"a", "b", "c", "d"}, make([]bool, 4), 0, 4)
}
