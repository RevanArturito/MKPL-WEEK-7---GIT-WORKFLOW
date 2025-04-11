package main

import "fmt"

func Substraction(a, b int) int {
	return a - b
}

// Branch Substraction
func main() {
	fmt.Println("Ini Branch Substraction")
	fmt.Println("hasil : ", Substraction(3, 4))
}
