package main

import "fmt"

func main() {
	x := []string{"a", "b", "c"}

	for v := range x {
		fmt.Println(v) //prints 0, 1, 2
	}
}
