func TestFunc() {
	u := &User{} // 포인터 변수를 선언하고 인스턴스를 생성한다. User라는 struct주소로 u를 만들었다.
	u.Age = 30   // (*u).age 라고 쓰는 것과 같다 -> u가 가리키고 있는 공간의 Age 필드 값을
	fmt.Println(u)
} // 내부 변수 u는 사라진다. 더불어 인스턴스도 사라진다.