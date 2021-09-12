package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // ServeMux 인스턴스 생성
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Go") // 인스턴스에 핸들러 등록
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	http.ListenAndServe(":3000", mux) // mux 인스턴스 사용

}

/**
Mux
multiplexer의 약자로 -> 여러 입력 중 하나를 선택해서 반환하는 디지털 장치를 말한다. 웹 서버는 각 URL에 해당하는 핸들러들을 등록한 다음
HTTP 요청이 왔을 때 URL에 해당하는 핸들러를 선택해서 실행하는 방시이다.
이 핸들러를 선택하고 실행하는 구조체 이름이 Mux를 제공한다고 해서 ServeMux이다. 비슷한 의미인 라우터 라고 말하기도 한다.
*/
