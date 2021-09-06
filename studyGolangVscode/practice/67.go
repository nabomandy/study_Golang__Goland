package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3} // len:3 cap:3 슬라이스 생성 -> slice1: [1,2,3] 3,3

	slice2 := append(slice1, 4, 5) // append() 함수로 요소를 추가 -> slice1: [1,2,3,4,5] 5,6

	fmt.Println("slice:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100 // slice1 요소값 변경 slice1: [1,100,3] slice2: [1,2,3,4,5]

	// 슬라이스 정보 출력
	fmt.Println("After change second element")
	fmt.Println("slice:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500) // slice1 요소값 변경

	// 슬라이스 정보 출력
	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

}
