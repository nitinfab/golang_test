package main

import (
	"fmt"
)

func main() {
	fmt.Println("Testing in Go")
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
