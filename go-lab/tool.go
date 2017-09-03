package main

import (
	"fmt"
)

func PrintIntegers(integers []int) {
	for _, i := range integers {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
