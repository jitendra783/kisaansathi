package utils

import (
	"math"
	"math/rand"
)

// Round... Perform rounding off upto certain decimal points using standard math library.
// for example:
//
//	Input: num = 3.14159265, precession = 2
//	Output: 3.14
func Round(num float64, precession uint) float64 {
	if precession == 0 {
		return math.Round(num)
	}
	multiperFactor := math.Pow(10, float64(precession))

	return math.Round(num*multiperFactor) / multiperFactor
}

// GenerateRandomNumber... Generate random number, takes number of digits required in output random number.
// for example:
//
//	Input: numDigit = 5
//	Output: 45896
func GenerateRandomNumber(numDigit int) int64 {
	var (
		max int64
		min int64
	)
	switch numDigit {
	case 2:
		min = 11
		max = 99
	case 3:
		min = 111
		max = 999
	default:
		return 0
	}

	return min + rand.Int63n(max-min)
}
