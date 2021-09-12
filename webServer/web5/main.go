package main

import (
	"awesomeProject/webServer/web5/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())

}
