package main

import "fmt"

const PI = 3.14              // 타입 없는 상수 -> 정하지 않았다는 의미
const FloatPI float64 = 3.14 // 타입을 정한 상수 -> float64 타입 상수 -> 특정 타입 값으로만 사용되는 것을 강제하고 싶을 때.

func main() {
	var a int = PI * 100          // 오류가 발생하지 않는다.
	var b int = FloatPI * 100     // 타입 오류 발생. 좌변은 int이고 우변은 float64이기 때문이다.
	var c float64 = FloatPI * 100 // 타입 오류 발생. 좌변은 int이고 우변은 float64이기 때문이다.
	// go에서는 양변의 타입이 같아야 연산이 가능하다!

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
