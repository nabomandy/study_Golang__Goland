type Element struct { // 구조체
	Value interface{} // 데이터를 저장하는 필드
    Next *Elelment // 다음 요소의 주소를 저장하는 필드
    Prev *Elelment // 이전 요소의 주소를 저장하는 필드