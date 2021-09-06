package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}

/*
1. Add 함수를 정의. 두 정수 타입 매개변수 a와 b를 입력으로 받아서
2. 그 둘의 합을 반환
3. Add() 함수를 호출. 반환값을 c에 저장.
4. c를 출력
*/
