package main

import (
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	assert.Equal(t, 82, square(9), "square(9) should be 81") // 아래 4줄이 assert.. 한줄로 대체
	//rst := square(9)
	//if rst != 81 {
	//	t.Errorf("square(() should be 81 but returns %d", rst)
	//}
}

func TestSquare2(t *testing.T) {
	assert.Equal(t, 9, square(3), "square(3) should be 9")
	//rst := square(3)
	//if rst !=9 {
	//	t.Errorf("square(9) should be 81 but result %d", rst)
	//}
}

func TestSquare3(t *testing.T) {
	rst := square(0)
	if rst != 0 {
		t.Errorf("square(0) should be 0 but returns %d", rst)
	}
}
