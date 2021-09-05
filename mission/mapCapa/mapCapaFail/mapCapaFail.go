package main

func main() {
	m := make(map[string]int, 99)
	cap(m) //error
}
