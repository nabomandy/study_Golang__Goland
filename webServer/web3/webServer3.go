package main

import (
	"awesomeProject/webServer/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())

}
