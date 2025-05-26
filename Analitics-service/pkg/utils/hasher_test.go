package utils_test

import (
	"fmt"
	"math/rand"
	"testing"
)

var str = "abcdefghijklmnopqrstuwyufyucnmOIFEWIUFPNWEYUBFRYBDIUJFR"

func GenerateString(length int) string {
	ans_str := ""
	for i := 0; i < length; i++ {
		rnd := rand.Intn(len(str))
		ans_str += string(str[rnd])
	}
	return ans_str
}

func BenchmarkStringContate(b *testing.B) {
	len := 10000
	first := GenerateString(len)
	second := GenerateString(len)
	// b.StartTimer()
	third := first + second
	// b.StopTimer()
	_ = third
}

func BenchmarkStringFormat(b *testing.B) {
	// b.StopTimer()
	len := 1000000
	first := GenerateString(len)
	second := GenerateString(len)
	// b.StartTimer()
	third := fmt.Sprintf("%s%s", first, second)
	// b.StopTimer()
	_ = third
}
