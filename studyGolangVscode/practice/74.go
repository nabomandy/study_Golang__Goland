package main

import "fmt"

type account struct {
	balance int
}

// 아래 두 func은 동작은 같다. 단지 의미상의 차이가 있다.
func withdrawFunc(a *account, amount int) { // 일반 함수 표현 -> 인자를 2개를 받고, 인자들 타입과는 상관이 없다
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) { //메서드 표현 -> 리시버가 정해져 있다 -> 함수가 리시버에 포함된 기능이라고 보면 된다.
	a.balance -= amount
}

func main() {
	a := &account{100} // balance가 100인 account 포인터 변수 생성 -> a의 타입은  포인트 타입, account 주소를 받으니까. 주소를 값으로 가지니까. *account

	withdrawFunc(a, 30) // 함수 형태 후출

	a.withdrawMethod(30) // 메서드 형태 호출 -> a라는 객체를 리시버로 받는 메서드를 호출

	fmt.Printf("%d\n", a.balance)
}
