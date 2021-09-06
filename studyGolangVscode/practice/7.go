package main

import "fmt"

func printNo(n int) {
	if n == 0 { // 재귀 호출 탈출 조건
		return
	}
	fmt.Println(n)
	printNo(n - 1)          // 재귀 호출
	fmt.Println("After", n) // 재귀 호출 이후 출력
}

func main() {
	printNo(3) // 함수 호출
}

/*
결과값
3
2
1
After 1
After 2
After 3
*/
