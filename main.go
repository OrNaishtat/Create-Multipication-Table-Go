package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Printf("%5v", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5v", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5v", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5v", i*j)

		}
		fmt.Println()
	}

}
