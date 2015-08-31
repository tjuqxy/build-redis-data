package tools

import (
	"time"
	"math/rand"
)

var (
	r *rand.Rand
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 包括传入的最大值, int值最小为1
func RandInt(maxInt int) int {
	return r.Intn(maxInt) + 1
}

func RandFloat(maxFloat int) float64 {
	return float64(maxFloat) * r.Float64()
}

func RandString(maxLen int) string {
	buff := make([]byte, 0)
	length := RandInt(maxLen)
	for {
		if length < 0 {
			break
		}
		buff = append(buff, byte(r.Intn(26) + 'a'))
		length--
	}
	return string(buff)
}

func RandRuneString(maxLen int) string {
	buff := make([]rune, 0)
	length := RandInt(maxLen)
	for {
		if length < 0 {
			break
		}
		buff = append(buff, rune(r.Intn(0xFF)))
		length--
	}
	return string(buff)
}
