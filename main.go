package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("How dare you...\n")

	fmt.Printf(fmt.Sprint("1 + 1 = ", add(1, 1), ", hahaha\n"))
}
