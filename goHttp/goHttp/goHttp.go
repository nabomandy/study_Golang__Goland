package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    // 쿼리 인수 가져오기
	name := values.Get("name") // 특정 키값이 있는지 확인
	if name == "" {
		name = "Go"
	}
	id, _ := strconv.Atoi(values.Get("id")) //id 값을 가져와서 int 타입 변환
	fmt.Fprintf(w, "Hello %s id:%d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":6000", nil)

}
