package main

import (
	"fmt"
)

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9}, // 이중 배열 사이에 ,로 구분한다. 마지막,를 찍는 거에 주의. 같은 줄에서 }}가 닫히면 안찍어도 되지만 이것처럼 }과 } 이 다른 줄에서 닫히면 찍어야한다.
	}

	for _, arr := range a { //배열 하나씩 끄집어 내서
		for _, v := range arr { // 배열을 다시 range를 돌아서 요소를 꺼낸다.
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

//이중 배열을 이중 for문을 써서 순회하는 것
