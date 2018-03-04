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

//Sumil returns the sum of all uint8 variables passed in
func Sumil(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

//Sumfl returns the sum of all float32 variables passed in
func Sumfl(a []float32) float32 {
	var sum float32
	for _, v := range a {
		sum += v
	}
	return sum
}

//Sumffl returns the sum of all float64 variables passed in
func Sumffl(a []float64) float64 {
	var sum float64
	for _, v := range a {
		sum += v
	}
	return sum
}

//Maxi finds the max of all values passed in
func Maxi(a ...int) int {
	cmax := a[0]
	for _, b := range a {
		if b > cmax {
			cmax = b
		}
	}
	return cmax
}

//Maxil finds the max of a list
func Maxil(a []int) int {
	cmax := a[0]
	for _, b := range a {
		if b > cmax {
			cmax = b
		}
	}
	return cmax
}

//Maxf finds the max of all values passed in
func Maxf(a ...float32) float32 {
	cmax := a[0]
	for _, b := range a {
		if b > cmax {
			cmax = b
		}
	}
	return cmax
}

//Maxfl finds the max of a list
func Maxfl(a []float32) float32 {
	cmax := a[0]
	for _, b := range a {
		if b > cmax {
			cmax = b
		}
	}
	return cmax
}

//Maxff finds the max of all values passed in
func Maxff(a ...float64) float64 {
	cmax := a[0]
	for _, b := range a {
		if b > cmax {
			cmax = b
		}
	}
	return cmax
}

//Maxffl finds the max of a list
func Maxffl(a []float64) float64 {
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

//Avgi returns the average of the values given
func Avgi(vals ...int) int {
	sum := Sumi(vals...)
	return sum / len(vals)
}

//Avgf returns the average of the values given
func Avgf(vals ...float32) float32 {
	sum := Sumf(vals...)
	return sum / float32(len(vals))
}

//Avgff returns the average of the values given
func Avgff(vals ...float64) float64 {
	sum := Sumff(vals...)
	return sum / float64(len(vals))
}

//Avgil returns the average of the values given
func Avgil(vals []int) int {
	sum := Sumil(vals)
	return sum / len(vals)
}

//Avgfl returns the average of the values given
func Avgfl(vals []float32) float32 {
	sum := Sumfl(vals)
	return sum / float32(len(vals))
}

//Avgffl returns the average of the values given
func Avgffl(vals []float64) float64 {
	sum := Sumffl(vals)
	return sum / float64(len(vals))
}
