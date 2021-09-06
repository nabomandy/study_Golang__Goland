//후처리 생략. 조건문만 있는 경우
// 무한루프 됨

package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i, ", ")
		i++ // 후처리를 생략하고 값이 변하고 싶으면 이렇게 적으면 된다.
	}
	fmt.Println() // 한 칸 띄우려는 목적
	fmt.Print(i)
}
