package main

import "fmt"

func main() {
	var a = 123
	var b int = 4567
	f := 3.14159269

	fmt.Printf("%6d\n, a")  // 출력 영역은 총 6칸. 따라서 %6d
	fmt.Printf("%06d\n, b") // 공란에 0이 채워지기 때문에 %06d
	fmt.Printf("%6d\n, c")  // 소수점 이하 2자리가 표시되기 때문에 %6.2f

}
