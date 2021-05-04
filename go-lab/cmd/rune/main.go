package main

import (
	"fmt"
)

func main() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Printf("%d. hello, world\n", i)
	//	}

	//	m0 := map[string]int{"a": 1, "b": 2, "c": 3}

	//	for k, v := range m0 {
	//		fmt.Printf("%s=>%d\n", k, v)
	//	}

	//	var s string = "hello"

	//	for i, c := range s {
	//		fmt.Printf("character[%d]: %c %s\n", i, c, string(c))
	//	}

	m1 := map[rune]int{'a': 1, 'b': 2, 'c': 3}

	var x, y int = 2, 5
	modifyMap(m1, 'c', x*y)
	printMap(m1)
	if checkMap(m1) {
		fmt.Println("check-map is true")
	}
}
func checkMap(m map[rune]int) bool {
	return m['c'] == 10
}

func modifyMap(m map[rune]int, c rune, i int) {
	m[c] = i
}

func printMap(m map[rune]int) {
	for c, v := range m {
		fmt.Printf("%c=>%d\n", c, v)
	}
}
