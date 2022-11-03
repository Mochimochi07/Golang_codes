package main

import (
	"fmt"
	"os"
)

func subtract(x, y int) int {
	return x - y
}

func add(x, y int) int {
	return x + y
}

func divide(x, y float32) float32 {
	return x / y
}

func multiply(x, y int) int {
	return x * y
}

func showoperators() {
	fmt.Println("What do you want to do? \n \n ")
	fmt.Println("1:add, 2:subtract, 3:divide, 4:multiply \n ")

}

func exitprogram() {

	fmt.Println("goodbye")
	os.Exit(5)

}

func main() {

	var a int
	var b int
	var choice string

	fmt.Println("Enter num1")
	fmt.Scanln(&a)
	fmt.Println("Enter num2")
	fmt.Scanln(&b)

	showoperators()
	fmt.Scanln(&choice)

	switch {

	case choice == "add" || choice == "1":

		fmt.Println(add(a, b))

	case choice == "subtract" || choice == "2":

		fmt.Println(subtract(a, b))

	case choice == "divide" || choice == "3":

		fmt.Println(divide(float32(a), float32(b)))

	case choice == "multiply" || choice == "4":

		fmt.Println(multiply(a, b))

	case choice == "exit":

		exitprogram()

	default:
		println("Not an operator please try again")
	}

}
