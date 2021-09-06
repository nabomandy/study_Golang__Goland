// 슬라이스와 배열의 동작 차이

package main

import "fmt"

func changeArray(array2 [5]int) { // 배열을 받아서 세 번째 값 변경
	array2[2] = 200
}

func changeSlice(slice2 []int) { // 슬라이스를 받아서 세 번째 값 변경
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array:", array)
	fmt.Println("slice", slice)

}

// 배열과 슬라이스의 구조가 다르기 때문에 결과가 다르게 나온다.

/*

결과값
array: [1 2 3 4 5]
slice [1 2 200 4 5]

*/
