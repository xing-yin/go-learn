package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/20
 * @Desc:
 */

func TestAbs(t *testing.T) {
	actual := Abs(-1)
	if actual != 1 {
		t.Errorf(" Abs(-1) = %f;expected 1", actual)
	}
}

func TestAbs_2(t *testing.T) {
	tests := []struct {
		x    float64
		want float64
	}{
		{-0.3, 0.3},
		{-2, 2},
		{-3.1, 3.1},
		{4, 4},
	}

	for _, tt := range tests {
		if got := Abs(tt.x); got != tt.want {
			t.Errorf("Abs() = %f,want %v", got, tt.want)
		}
	}
}

func TestAbs_3(t *testing.T) {
	tests := []struct {
		x    float64
		want float64
	}{
		{-0.3, 0.3},
		{-2, 2},
		{-3.1, 3.1},
		{4, 4},
	}

	for _, tt := range tests {
		got := Abs(tt.x)
		assert.Equal(t, got, tt.want)
	}
}

func TestMax(t *testing.T) {
	actual := Max(1, 2)
	if actual != 2 {
		t.Errorf("Max(1,2) = %f;expected 1", actual)
	}
}

func TestMin(t *testing.T) {

}

func TestRandInt(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandInt(); got != tt.want {
				t.Errorf("RandInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
