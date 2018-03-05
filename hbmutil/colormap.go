package hbmutil

import "fmt"

//Gscmap returns a float32 N by 3 matrix with resultant r,g,b values for (vals) discrete levels. Similar to Matlab's 'gray()' command.
//Thus, Gscmap will return a (12*n) byte matrix.
func Gscmap(vals, max int) *Matrixf {
	resultantcmap := *NewMatrixf(vals, 3) //create the n x 3 colormap rgb matrix to be returned and dereference it.
	level := 1.0 / (float32(vals) - 1)    //determine the increments necessary to create a clean transition between inclusive 0 and 1
	for i := 0; i < vals; i++ {           //iterate through each value of the colormap
		resultantcmap.Data[vals-i-1] = []float32{level * float32(i), level * float32(i), level * float32(i)} // since this colormap is greyscale, all values will be the same per row.
	}
	resultantcmap = resultantcmap.Scale(float32(max)) //Finally, scalar multiply each element by 255 so that the resulting array fits the standard 32B dynamic range [0,255]
	return &resultantcmap                             //return the pointer to the new matrix
}

//Hotmap returns a float32 N by 3 matrix with resultant r,g,b values for (vals) discrete levels. Similar to Matlab's 'hot()' command.
//Thus, Hotmap will return a (12*n) byte matrix.
func Hotmap(vals, max int) *Matrixf {
	fmt.Println(vals, max)
	resultantcmap := *NewMatrixf(vals, 3) //create the n x 3 colormap rgb matrix to be returned and dereference it.

	level := 1.0 / (float32(vals / 3)) //determine the increments necessary to create a clean transition between inclusive 0 and 1
	level2 := 1.0 / (float32(vals/3) - float32(max)/4)
	for i := 0; i < vals/3; i++ { //iterate through each value of the colormap
		resultantcmap.Data[i] = []float32{level * float32(i), 0, 0}
	}
	for i := vals / 3; i < 2*vals/3; i++ { //iterate through each value of the colormap
		resultantcmap.Data[i] = []float32{1, level * float32(i-vals/3), 0}
	}
	for i := 2 * vals / 3; i < vals-1; i++ { //iterate through each value of the colormap
		resultantcmap.Data[i] = []float32{1, 1, level2 * float32(i-2*vals/3)}
	}
	resultantcmap.Data[vals-1] = []float32{1, 1, 1}
	resultantcmap = resultantcmap.Scale(float32(max)) //Finally, scalar multiply each element by 255 so that the resulting array fits the standard 32B dynamic range [0,255]
	return &resultantcmap                             //return the pointer to the new matrix
}
