package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10       //myInt 타입 변수
	fmt.Println(a.add(30)) // myInt타입의 add() 메서드 호출
	var b int = 20
	fmt.Println(myInt(b).add(50)) //int 타입을 타입 변환

}
