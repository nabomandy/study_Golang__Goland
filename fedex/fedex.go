// Fedex에서 제공한 패키지
package fedex

import (
	"fmt"
	_ "fmt"
)

// Fedex에서 제공한 패키지 내 전송을 담당하는 구조체이다
type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)

}

/**
추상화란
내부 동작(구현)을 감춰서 서비스 제공자와 사용자 모두에게 자유를 주는 방식

추상화를 하면 의존성을 끊을 수 있다.

인터페이스 -> 디커플링 -> 의존성을 제거하는 데 가장 강력한 도구

추상계층 (Abstract layer)//
서비스 제공자가 있고 서비스 사용자가 있다. -> 제공자와 사용자간의 상호작용을 정의하는 거
인터페이스는 이 둘간에 추상계층을 제공한다

덕타이핑 -> 잘 아셔야 해요ㅇㅁ
만약 어떤 새를 봤는데 그 새가 오리처럼 걷도 오리처럼 날고 오리처럼 소리내면
나는 그 새를 오리라고 부르겠다.
타입 선언시 인터페이스 구현 여부를 명시하지 않아도 된다.

덕타이핑 -> 사용자 중심의 코딩이 된다.

인터페이스 구현여부를 타입선언시 하지 않고, 사용될 때 결정하기 때문에
서비스 제공자는 구체화된 객체(Concrete class)를 제공하고 사용자가 필요에 따라
인터페이스를 정의해서 사용할 수 있다.

서비스 제공자는 인터페이스를 제공할 필요가 없고, 구현객체만, 타입만 제공하면 된다.
사용자 입장에서는 구체화된 객체를 그냥 쓰면된다.

구현체를 쓰다가 나중에 인터페이스를 만들어도 된다.
미리 만들어도 좋지만, 낭비일 수 있다. 필요가 생기면 인터페이스르 만들어서 바꿔도 된다. -> 덕타이핑이기 때문에 가능하다.
자바 같은 경우에는 미리 만들어야 된다.

덕타이핑의 장점
단점-> 잘 모르겠다.

OOP, c++, JAVA 이 때는 덕타이핑 개념이 없었다. 파이썬이 나오면서 생겼고, 고도 가져온거.
언어 자체적으로 덕타이핑 자체의 구현 난이도가 올라간다.

자바같은 경우는 클래스를 만들 때 인터페이스를 염두에 둬야 한다.
OOP로 가면서 설계가 중요해졌다. -> 클래스간의 관계설정 -> 관계가 바로 인터페이스
자바의 방식은 설계 에서 구현으로 이어지는 방식 -> 포커스 방식(전통적 방식)

고에서는 인터페이스를 미리 뽑아낼 필요가 없다. 구현을 할 때 설계에서 나온 인터페이스 중심으로 할 필요는 없다.
고 -> 설계는 항상 바뀔 수 있다. 구현도 바뀔 수 있다.(에자일 방식, 스크럼 방식에 더 잘 어울린다.)


서로 다른 부분은 어댑터를 사용해서 맞춰줘야 한다. -> 어댑터 패턴
덕타이핑의 사용여부 상관없이 맞추는 과정은 대게  필요하다.

*/
