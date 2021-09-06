package main

import "fmt"

func main() {
	str1 := "Hello\t'World'\n" // 큰 따옴표로 묶으면 특수 문자가 동작한다.
	str2 := `Go is "awesome"!\nGo is simple and\t'powerful'`

	fmt.Println(str1)
	fmt.Println(str2)

}

/*
결과값
Hello	'World'

Go is "awesome"!\nGo is simple and\t'powerful'
*/
