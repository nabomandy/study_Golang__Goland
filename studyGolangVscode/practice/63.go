package main

import "fmt"

func main() {
	var slice []int

	for i := 1; i <= 10; i++ { // 요소를 하나씩 추가
		slice = append(slice, i)
	}

	slice = append(slice, 11, 12, 13, 14, 15) // 한번에 여러 요소 추가
	fmt.Println(slice)

}

/*

결과값
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]

*/
