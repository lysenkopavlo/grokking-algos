package main

import "fmt"

func countdown(i int) {
	fmt.Println(i)
	if i <= 0 {
		return
	}

	countdown(i - 1)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	countdown(5)
	fmt.Println(factorial(5))

}
