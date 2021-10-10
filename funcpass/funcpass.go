package main

func main() {
	msg := "Hello"
	say(&msg)
	println(msg) //변경된 메시지 출력
}

func say(msg *string) {
	println(*msg)
	*msg = "Changed" //메시지 변경
}
