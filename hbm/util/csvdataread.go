package hbmutil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

//ReadMEPcsvf reads MEPs out of a comma separated values file, ignoring non-numeric characters.
func ReadMEPcsvf(file string) *Matrixf {
	csvfile, err := os.Open(file) // open the csv fiel

	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close() // close the file to prevent memory leaks once we are done with it
	MEPrunes := make([]rune, 0, 10000)
	reader := bufio.NewReader(csvfile)
	for i := 0; ; i++ {
		rn, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		MEPrunes = append(MEPrunes, rn)
	}
	fields := 0
	for i := 0; i < len(MEPrunes); i++ {
		if MEPrunes[i] == 13 || MEPrunes[i] == 44 {
			fields++
		}
	}
	out := NewMatrixf(fields/4+1, 4)
	col := 0
	row := 0
	curval := ""
	for i := 0; i < len(MEPrunes); i++ {
		if MEPrunes[i] == 46 {
			curval += "."
		} else if MEPrunes[i] > 47 && MEPrunes[i] < 58 {
			curval += string(MEPrunes[i])
		} else if MEPrunes[i] == 13 || MEPrunes[i] == 44 {
			val, _ := strconv.ParseFloat(curval, 32)
			out.Data[row][col] = float32(val)
			if col == 3 {
				col = 0
				row++
			} else {
				col++
			}
			curval = ""
		}
	}
	if out.Data[out.Rows-1][0] == 0 && out.Data[out.Rows-1][1] == 0 && out.Data[out.Rows-1][2] == 0 && out.Data[out.Rows-1][3] == 0 {
		out.Data = out.Data[:out.Rows-1]
		out.Rows--
	}
	return out
}

//parseptstoptsf takes a 2-dimensional array and returns an array of hbmutil.Point3's.
func parseptstoptsf(data [][]string, xcol, ycol, zcol, MEPcol int) []Point3f {
	out := make([]Point3f, len(data))
	for j := 0; j < len(data); j++ { // for each row (separate entry)
		x, errx := (strconv.ParseFloat(data[j][xcol], 32))   // parse for a float in the x column
		y, erry := (strconv.ParseFloat(data[j][1], 32))      // parse for a float in the y column
		z, errz := (strconv.ParseFloat(data[j][2], 32))      // parse for a float in the z column
		m, errm := (strconv.ParseFloat(data[j][MEPcol], 32)) // parse for a float in the intensity column
		if errx != nil || erry != nil || errz != nil {
			fmt.Println("Error(s) converting coordinates: ", errx, erry, errz)
		}
		if errm != nil {
			fmt.Println("Error converting MEP value: ", errm)
		}
		out[j] = Point3f{float32(x), float32(y), float32(z), float32(m)}
	}
	return out
}

//parseptstomatf takes a 2-dimensional array and returns an n x 4 matrix of points with intensity values (x,y,z,intensity)
func parseptstomatf(data [][]string, xcol, ycol, zcol, MEPcol int) *Matrixf {
	out := NewMatrixf(len(data), 4)  // make a new matrix to hold the points
	for j := 0; j < len(data); j++ { // for each row (separate entry)
		fmt.Println(data[j][:])
		x, errx := (strconv.ParseFloat(data[j][xcol], 32))   // parse for a float in the x column
		y, erry := (strconv.ParseFloat(data[j][ycol], 32))   // parse for a float in the y column
		z, errz := (strconv.ParseFloat(data[j][zcol], 32))   // parse for a float in the z column
		m, errm := (strconv.ParseFloat(data[j][MEPcol], 32)) // parse for a float in the intensity column
		if errx != nil || erry != nil || errz != nil {
			fmt.Println("Error(s) converting coordinates: ", errx, erry, errz)
		}
		if errm != nil {
			fmt.Println("Error converting MEP value: ", errm)
		}

		out.Unpackr(j, float32(x), float32(y), float32(z), float32(m)) // unpack the x,y,z,intensity values into the corresponding row of our new matrix
	}
	return out
}

//MattoImgi converts a list of points with intensity values to an image array
func MattoImgi(pts *Matrixi, xcol, ycol, wid, hei int) *Matrixi {
	target := NewMatrixi(256, 256)  //Create an empty matrix for us to store the image information in
	for j := 0; j < pts.Rows; j++ { // for each point,
		x := pts.Data[j][xcol]             // take the x point
		y := pts.Data[j][ycol]             // take the y point
		target.Data[y][x] = pts.Data[j][3] // and put the intensity value at the corresponding x and y value in the new matrix
	}
	return target
}

//MattoImgf converts a list of points with intensity values to an image array
func MattoImgf(pts *Matrixf, xcol, ycol, wid, hei int) *Matrixf {
	target := NewMatrixf(256, 256)  //Create an empty matrix for us to store the image information in
	for j := 0; j < pts.Rows; j++ { // for each point,
		x := int(Roundf(pts.Data[j][xcol])) // take the x point
		y := int(Roundf(pts.Data[j][ycol])) // take the y point
		target.Data[y][x] = pts.Data[j][3]  // and put the intensity value at the corresponding x and y value in the new matrix
	}
	return target
}

//MattoImgff converts a list of points with intensity values to an image array
func MattoImgff(pts *Matrixff, xcol, ycol, wid, hei int) *Matrixff {
	target := NewMatrixff(256, 256) //Create an empty matrix for us to store the image information in
	for j := 0; j < pts.Rows; j++ { // for each point,
		x := int(Roundff(pts.Data[j][xcol])) // take the x point
		y := int(Roundff(pts.Data[j][ycol])) // take the y point
		target.Data[y][x] = pts.Data[j][3]   // and put the intensity value at the corresponding x and y value in the new matrix
	}
	return target
}
