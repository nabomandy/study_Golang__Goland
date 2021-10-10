package myapp

import (
	"fmt"
	mux2 "github.com/gorilla/mux"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")

}

func userHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "Get User info by /users/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux2.Vars(r)
	fmt.Fprint(w, "User Id:", vars["id"])

}

// NewHandler make a new myapple Handler
func NewHandler() http.Handler {
	mux := mux2.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", userHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

	return mux

}
