package main

import "fmt"

func factorial(n float64) float64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n, e, i float64
	fmt.Scan(&n)
	for i = 0; i <= n; i++ {
		e += 1 / factorial(i)
	}
	fmt.Println(e)
}
