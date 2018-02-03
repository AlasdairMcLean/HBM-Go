package hbmutil

//Gscmap returns an N by 3 matrix with resultant r,g,b values for (vals) discrete levels. Similar to Matlab's 'gray()' command.
func Gscmap(vals int) *Matrixf {
	resultantcmap := *NewMatrixf(vals, 3)
	var maxval float32 = 1.0
	level := maxval / (float32(vals) - 1)
	for i := 0; i < vals; i++ {
		resultantcmap.Data[vals-i-1] = []float32{level * float32(i), level * float32(i), level * float32(i)}
	}
	resultantcmap = resultantcmap.STimes(float32(255))
	return &resultantcmap
}
