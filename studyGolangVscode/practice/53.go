package main

import "fmt"

func main() {
	str := "Hello World"

	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100} // 'H', 'e', 'l', 'l', 'o', 'W', 'o', 'r', 'l', 'd' 문자코드 배열

	fmt.Println(str)
	fmt.Println(string(runes))
}
