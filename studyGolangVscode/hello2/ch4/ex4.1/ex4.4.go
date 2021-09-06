package main

import "fmt"

func main() {
	a := 3              // int
	var b float64 = 3.5 // float 64

	var c int = int(b)  // float64에서 int로 변환
	d := float64(a * c) // int에서 float로 변환

	var e int64 = 8
	f := int64(d) * e // float64에서 int 64로 변환

	var g int = int(b * 3) // float64에서 int로 변환
	var h int = int(b) * 3 // float64에서 int로 변환. g와 값이 다르다.
	fmt.Println(g, h, f)

}
