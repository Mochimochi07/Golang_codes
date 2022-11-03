package main

import "fmt"

func divide(x, y float32) float32 {
	return x / y
}

func main() {

	var a float32
	var b float32
	fmt.Println("Enter num1")
	fmt.Scanln(&a)
	fmt.Println("Enter num2")
	fmt.Scanln(&b)
	fmt.Println("The quotient of two given numbers")
	fmt.Println(divide(a, b))

}
