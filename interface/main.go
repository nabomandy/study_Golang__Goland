package main

import (
	"fmt"
	_ "fmt"
)

/***
Go 언어에서는 -er을 붙여서 인터페이스명을 만드는 것을 권장하고 있습니다. String() 메서드를 가진 인터페이스란 뜻으로 Stringer 라고 만들었다.
Stringer의 원뜻과는 관계가 없다.
*/

type Stringer interface { // Stringer 인터페이스 선언
	String() string // 매개변수 없이 string 타입을 반환하는 String()메서드를 포함
	//String1() string // 이렇게 하면 컴파일 에러 -> String1()이라는 메서드가 없기 떄문에.
	//인터페이스 구현여부를 타입 선언시 하는게 아니라 인터페이스에 대입을 할 때 체크하는 거
	// -> 여기 있는 인터페이스를 가지고 있으면, Stringer 인터페이스라고 보겠다.
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string { // Student의 String() 메서드

	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name) // 문자열 만들기 Sprintf 와 printf 의 차이 확인해볼 것
}

func main() {
	student := Student{"앤디", 31} // Student 타입
	var stringer Stringer        // Stringer 타입

	stringer = student // stringer 값으로 student 대입

	fmt.Printf("%s\n", stringer.String()) // stringer 의 String() 메서드 호출

}

/**
인터페이스
구체화된 객체(Concrete object)가 아닌 추사화된 상호작용으로 관계를 표현

메서드는 결국 관계를 나타내는 것이다.

OOP -> 관계 중심
어떤 클래스가 어떤 관계를 맺고 있느냐 -> 프로그램 설계
고에서는 메서드로 이 관계를 나타낸다.
메서드는 구현이 포함되어 있다. 구체적인 코드가 들어가 있다.
메서드는 관계 + 구현인데, -> 여기서 구현부분을 빼고 관계만 나타내는 게 인터페이스 이다.
추상화된 상호작용으로 관계를 표현하는게 인터페이스이다.

인터페이스도 타입이다 -> 변수를 선언할 수 있다 -> 인터페이스 타입의 변수를 만들 수 있다.
인터페이스는 선언할 떄, {} 안에 구현이 빠진 메서드들을 넣는다.

어떤 객체가 인터페이스를 구현하기 위해서는 { }안의 메서드를 가지고 있어야 한다. -> 어떠한 A가 { } 안의 메서드를 포함하고 있으면,
해당 인터페이스를 구현하고 있다고 보는 것이다.

* 메서드는 반드시 메서드명이 있어야 한다.
* 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없다.
* 인터페이스에서는 메서드 구현을 포함하지 않는다.


*/
