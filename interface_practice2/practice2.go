// 다음 패킺지의 OurDB 구조체의 모든 공개된 메서드를 이용하는 인터페이스르 만들어보세요ㅋ

package AwesomeDB

type DataBase interface {
	GetData() string
	WriteData(data string)
	Close() error
}

type OurDB struct {
	Name string
}

func (db *OurDB) GetData() string {
	...
}

func (db OurDB) WriteData(data string) {
	...
}

func (db OurDB) Close() error{
	...
}

func main() {

}
