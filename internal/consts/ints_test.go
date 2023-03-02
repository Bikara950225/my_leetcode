package consts

import (
	"math"
	"testing"
)

func TestIntConst(t *testing.T) {
	if math.MinInt64 != MinInt64 {
		t.Errorf("TestIntConst error, MinInt64: %d != %d", MinInt64, math.MinInt64)
		return
	}
	if math.MaxInt64 != MaxInt64 {
		t.Errorf("TestIntConst error, MaxInt64: %d != %d", MaxInt64, math.MaxInt64)
		return
	}
}
