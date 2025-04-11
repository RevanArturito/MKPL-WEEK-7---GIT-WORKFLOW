package main

import "fmt"


func Substraction(a, b int) int {
	return a - b
}

func Addition(a, b int) int {
	return a + b
}

// Branch Main
func main() {
	fmt.Println("Ini Branch Substraction")
	fmt.Println("hasil : ", Substraction(3, 4))

	fmt.Println("Ini Branch Addition")
	fmt.Println("hasil : ", Addition(3, 4))

}
