package rand

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxMask = (1 << 6) - 1
)

// RandomString return specified length string
func RandomString(l int) string {
	b := make([]byte, l)
	s := len(letterBytes)
	for i := 0; i < l; {
		if idx := int(rand.Int63() & letterIdxMask); idx < s {
			b[i] = letterBytes[idx]
			i++
		}
	}

	return string(b)
}
