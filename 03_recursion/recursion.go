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

func euclid(a, b int) int {
	if b == 0 {
		return a
	} else {
		return euclid(b, a%b)
	}
}

// hanoi recursively prints the solution of hanoi tower problem
// i - starting column, k - finish column, n - number of disks
func hanoi(i, k, n int) {
	if n == 1 {
		fmt.Printf("Move disk 1 from column %d to column %d\n", i, k)
	} else {
		tmp := 6 - i - k // tmp - the third, temporary column
		hanoi(i, tmp, n-1)
		fmt.Printf("Move disk %d from column %d to column %d\n", n, i, k)
		hanoi(tmp, k, n-1)
	}
}

// binGenerator recursively generates every 1/0 combination in the length of n
func binGenerator(n int) {
	comb := make([]int, n)
	top := 0
	if n != 0 {
		comb[top] = 1
		top++
		binGenerator(n - 1)
		top--

		comb[top] = 0
		top++
		binGenerator(n - 1)
		top--
	} else {
		for i := 0; i < top; i++ {
			fmt.Println(comb[i])
		}
	}
}

func main() {
	countdown(5)
	fmt.Println("Factorial:", factorial(5))

	fmt.Println("Euclid:", euclid(4561, 987))

	hanoi(1, 2, 4)

	binGenerator(3)
}
