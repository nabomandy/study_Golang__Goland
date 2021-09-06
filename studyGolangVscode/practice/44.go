// 44, 45 함께 볼 것

package main

import "fmt"

type Data struct { //Data 타입 구조체
	value int
	data  [200]int
}

func ChangeData(arg Data) { // 매개변수로 Data를 받는다. 파라미터 이름은 arg이고 타입은 Data. 함수에 인자로 어떤 변수가 쓰이면 무조건 r-value -> 값으로 쓰인다는 의미
	arg.value = 999
	arg.data[100] = 999

}

func main() {
	var data Data

	ChangeData(data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100]= %d\n", data.data[100])

}
