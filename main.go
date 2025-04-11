package main

import "fmt"

func Addition(a, b int) int {
	return a + b
}

// Branch Addition
func main() {
	fmt.Println("Ini Branch Addition")
	fmt.Println("hasil : ", Addition(3, 4))
}
