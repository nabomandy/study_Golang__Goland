package main

type ReadWriter interface {
	Read()
	Write()
}

type File struct {
}

func (f *File) Read() { // Write()메서드는 없음
}

//func (f *File) Write() {
//}

func ReadWrite(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	f := &File{}
	ReadWrite(f)

}

/**

File 타입은 Read() 메서드만 가지고 있어 Read()와 Write()가 정의된 ReadWriter 인터페이스로 사용할 수 없어 에러가 발생한다.

*/
