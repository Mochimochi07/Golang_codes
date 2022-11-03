package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {

	var a int
	var b int
	fmt.Println("Enter num1")
	fmt.Scanln(&a)
	fmt.Println("Enter num2")
	fmt.Scanln(&b)
	fmt.Println("The sum of two given numbers")
	fmt.Println(add(a, b))

}
