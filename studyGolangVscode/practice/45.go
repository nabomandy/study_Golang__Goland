package main

import "fmt"

type Data struct { //Data 타입 구조체
	value int
	data  [200]int
}

func ChangeData(arg *Data) { // 매개변수로 Data를 받는다. 파라미터 이름은 arg이고 타입은 Data. 함수에 인자로 어떤 변수가 쓰이면 무조건 r-value -> 값으로 쓰인다는 의미

	// ChagedData(arg *Data) -> Data type의 주소를 값으로 가질 수 있다.
	// ChagedData(arg &Data) -> Data type의 주소 값

	arg.value = 999
	arg.data[100] = 999

}

func main() {
	var data Data

	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100]= %d\n", data.data[100])

}

/*
var data Data
var arg *Data
arg = &data // data의 주소 값으로 정한다

arg.value = 999
위 얘기는
(*arg).value = 999 와 같다 -> *을 붙이면 공간 -> 공간에 포함된 값을 999로 바꿔라


포인터는 항상 주소를 값으로 가진다. 이걸 기억해야 한다!

포인터를 썼을 때 두가지 의미가 있는데
1. 같은 공간을 가리켜서 그 공간의 값을 바꿀 수 있다. 함수내부에서
2. 복사되는 사이즈가 메모리 주소값만 복사된다.

*/
