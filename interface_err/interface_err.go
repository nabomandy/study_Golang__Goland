package main

type Stringer interface {
	String() string
}
type Student struct {
}

func main() {
	var stringer Stringer
	student := stringer.(*Student)

	/** 위 코드는 불가능하다. Student가 Stringer() 메서드를 가지고 잇지 않기 때문

	컴파일 에러 : stringer.(*Student) evaluated but not use
	런타임 에러 :
	# awesomeProject/interface_err
	.\interface_err.go:11:10: impossible type assertion:
		*Student does not implement Stringer (missing String method)

	Compilation finished with exit code 2
	*/

}
