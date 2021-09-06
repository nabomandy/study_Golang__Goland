// 슬라이싱 기능을 사용해 배열의 일부를 나타내는 슬라이스를 만드는 예제를 사용해 자세히 살펴보자.

package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2] // 슬라이싱

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100 // array의 두 번째 값 변경

	fmt.Println("After change second  element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500) // slice에 값 추가

	fmt.Println("After append 500")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

}
