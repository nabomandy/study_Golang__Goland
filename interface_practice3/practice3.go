package main

type Stringer interface {
	String()
}

type Reader interface {
	Read()
}

func CheckANdRun(stringer Stringer) {
	r := stringer.(Reader)
	r.Read()

}

func main() {

}
