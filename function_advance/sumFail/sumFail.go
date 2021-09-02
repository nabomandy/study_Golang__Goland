package main

import (
	"fmt"
	_ "fmt"
)

// nums는 []int{1,2,3,4,5}와 같다 -> 가변인수 -> 함수 내부에서 슬라이스 타입으로 동작함을 기억!
func sum(nums int) int { // 가변 인수를 받는 함수
	sum := 0

	fmt.Printf("nums 타입:%T\n", nums) // nums 타입 출력

	for _, v := range nums { // 컴파일 에러 : Cannot range over 'nums' (type int)
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // 인수 5개를 사용. 컴파일 에러 : Too many arguments in call to 'sum'
	fmt.Println(sum(10, 20))        // 인수 2개를 사용. 컴파일 에러 : Too many arguments in call to 'sum'
	fmt.Println(sum())              // 인수 0개를 사용. 컴파일 에러 : Not enough arguments in call to 'sum'

}
