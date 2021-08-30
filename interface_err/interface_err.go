package main

type Stringer interface {
	String() string
}
type Student struct {
}

func main() {
	var stringer Stringer
	stringer.(*Student) // 불가능하다. Student가 Stringer() 메서드를 가지고 잇지 않기 때문

}
