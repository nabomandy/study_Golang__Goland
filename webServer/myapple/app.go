package myapple

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    // 쿼리 인수 가저욕괴
	name := values.Get("name") // 특정 키값이 있는지 확인
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id")) //id 값을 가져와서 int 타입 변환
	fmt.Fprintf(w, "Hello %s! id: %d", name, id)
}

type fooHandler struct {
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User) // 값을 채워줄 인스턴스를 먼저 만들고
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now() // user의 CreatedAt 값을 현재시간으로 바꿔줄거에요 user 데이터를 성공적으로 디코드 하고 온 거니까요

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK) // 잘했다고 알려주고
	fmt.Fprint(w, string(data))  //json은 스트링 형태인데 data는 바이트 형태기 때문이

}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", barHandler) // bar 핸들러 등록

	mux.Handle("/foo", &fooHandler{}) // 인터페이스 형태로 사용하는 방법
	return mux
}
