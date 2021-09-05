package main


func TotalTime(db Database) int { // 컴파일 에러 : Unresolved type 'Database' Database 에 대한 타입 선언이 없다.
	db.Get
	db.Set
}


func Compare() {
	BDB := &BDatabase{}
	CDB := &CDatabase{}
}

if TotalTime(BDB) < TotalTime(CDB){
	fmt.q
}

func main() {

}
