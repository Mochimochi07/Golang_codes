package main

import (
	"fmt"
	"time"
)

func telltime() {

	fmt.Println("The time is", time.Now())

}

func main() {
	fmt.Println("Tell the time")

	fmt.Println("The time is", time.Now())

	telltime()
}
