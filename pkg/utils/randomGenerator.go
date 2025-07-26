package utils

import (
	"math/rand"
	"time"
)

const letterBytes = "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqWrRsStTuUvVwWxXyYzZ01234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

//generates Random string of length provided as argument
//	string created with combination of a-z, A-Z, 0-9
//	input:
// 		length of string
//	output:
// 		string
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[seededRand.Intn(len(letterBytes))]
	}
	return string(b)
}
