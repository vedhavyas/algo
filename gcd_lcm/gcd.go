package main

import "fmt"

func GCD(a, b int) int {
	if a == b {
		return a
	}

	n, d := a, b
	if b > a {
		n, d = b, a
	}

	r := n % d
	if r == 0 {
		return d
	}

	return GCD(r, d)
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func main() {
	fmt.Println("GCD(24, 36):", GCD(24, 36))
	fmt.Println("LCM(24, 36): ", LCM(24, 36))
}
