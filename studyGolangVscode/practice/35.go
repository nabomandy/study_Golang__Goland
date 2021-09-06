package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t { // range는 키워드 이다. range다음에 변수가 오는데,
		// 자료구조슬라이스 스트링 맵 등등
		// range는 for문 안쪽에서 쓰이는 키워드인데.
		// range + 자료구조가 올 경우에 각 요소들을 순회해준다.
		// i, v는 변수명으로 정해진 건 아니지만 관습적으로 저렇게 많이 쓴다.
		// 선언 해서 사용하지 않을 때는 에러가 난다. 그래서 _로 처리해줘야한다.
		fmt.Println(i, v)
	}
}
