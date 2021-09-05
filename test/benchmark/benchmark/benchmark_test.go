package main

import "testing"

func Test_fibonacci1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci1(tt.args.n); got != tt.want {
				t.Errorf("fibonacci1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonacci2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci2(tt.args.n); got != tt.want {
				t.Errorf("fibonacci2() = %v, want %v", got, tt.want)
			}
		})
	}
}
