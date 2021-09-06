// 방에 불을 키는 비트플래그 예제

package main

import "fmt"

const (
	MasterRoom uint8 = 1 << iota // 1
	LivingRoom                   // 2
	BathRoom                     // 4
	SmallRoom                    // 8
)

// 비트플래그에서 기본적으로 쓰는 유형
func SetLight(rooms, room uint8) uint8 {
	return rooms | room // 합쳐진다.
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room //비트클리어 연산자 -> 해당 비트값을 0으로 만든다.
}

func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room // 0인애는 무조건 0. 두개의 값이 1이면 1
}

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 불을 켠다")
	}

	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불을 켠다")
	}

	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 불을 켠다")
	}

	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은방에 불을 켠다")
	}

}

func main() {
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)

	rooms = ResetLight(rooms, SmallRoom)

	TurnLights(rooms)
}
