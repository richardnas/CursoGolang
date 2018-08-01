package main

import (
	"fmt"
)

func main() {

	fmt.Println("For for:")
	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}

	fmt.Println("For while:")
	y := 6
	for y <= 10 {
		fmt.Println(y)
		y = y + 1
	}

}
