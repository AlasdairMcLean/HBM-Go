//Math.go holds several useful mathematical functions for convenience

package hbmutil

import "math"

//Sumi returns the sum of all uint8 variables passed in
func Sumi(a ...int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

//Sumf returns the sum of all float32 variables passed in
func Sumf(a ...float32) float32 {
	var sum float32
	for _, v := range a {
		sum += v
	}
	return sum
}

//Sumff returns the sum of all float64 variables passed in
func Sumff(a ...float64) float64 {
	var sum float64
	for _, v := range a {
		sum += v
	}
	return sum
}

//Max finds the max
func Max(a []float64) float64 {
	cmax := a[0]
	for _, b := range a {
		if b > cmax {
			cmax = b
		}
	}
	return cmax
}

//Roundf returns the rounded value of input
func Roundf(f float32) float32 {
	return float32(math.Floor(float64(f) + .5))
}

//Roundff returns the rounded value of input
func Roundff(f float64) float64 {
	return math.Floor(f + .5)
}
