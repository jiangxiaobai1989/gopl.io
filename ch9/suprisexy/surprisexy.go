package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y int
	fmt.Print("x:", x, " ", "y:", y, "\n")
	go func() {
		x = 1
		fmt.Print("go_y:", y, " ", "\n") //A2
	}()

	fmt.Print("x:", x, " ", "y:", y, "\n")

	go func() {
		y = 1                            // B1
		fmt.Print("go_x:", x, " ", "\n") // B2
	}()

	time.Sleep(3)
	fmt.Print("x:", x, " ", "y:", y, "\n")
}
