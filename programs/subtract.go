package main

import "fmt"

func subtract(x, y int) int {
	return x - y
}

func main() {

	var a int
	var b int
	fmt.Println("Enter num1")
	fmt.Scanln(&a)
	fmt.Println("Enter num2")
	fmt.Scanln(&b)
	fmt.Println("The diff of two given numbers")
	fmt.Println(subtract(a, b))

}
