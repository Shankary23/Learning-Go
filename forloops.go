package main

import "fmt"

func runforloop() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}

func runwhileloop() {
	counter := 0
	sum := 0
	for counter < 10 {
		fmt.Println(counter)
		sum += counter
		counter += 1
	}
	fmt.Println(sum)
}

func recursion(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + recursion(n-1)
}
