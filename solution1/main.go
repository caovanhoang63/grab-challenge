package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) bool {
	oddCount := 0
	evenCount := 0
	// Implement your solution here
	for i := range A {
		if A[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return oddCount == evenCount
}

func main() {
	test1 := []int{-1, -3, 4, 7, 7, 7}
	if Solution(test1) {
		fmt.Println("Test1 passed")
	} else {
		fmt.Println("Test1 failed")
	}
}
