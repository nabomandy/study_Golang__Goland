//ch29/ex29.1/ex29.1.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World") // 웹 핸들러 등록
	})

	http.ListenAndServe(":3000", nil) // 웹서버시작

}
