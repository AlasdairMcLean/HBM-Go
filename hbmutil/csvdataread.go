package hbmutil

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ReadMEP takes MEP values from a csvfile and returns an array of hotspot points.
func ReadMEP(filename string, MEPcol int) (*Matrixf, error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	file := strings.Join([]string{path, filename}, "/")
	fmt.Println(file)
	csvfile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	reader.LazyQuotes = true
	MEPs, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	MEPpoints := parsepoints(MEPs, MEPcol)
	return MEPpoints, nil
}

func parsepoints(data [][]string, MEPcol int) *Matrixf {
	out := NewMatrixf(len(data), 4)
	for j := 0; j < len(data); j++ {
		x, errx := (strconv.ParseFloat(data[j][0], 32))
		y, erry := (strconv.ParseFloat(data[j][1], 32))
		z, errz := (strconv.ParseFloat(data[j][2], 32))
		m, errm := (strconv.ParseFloat(data[j][MEPcol-1], 32))
		if errx != nil || erry != nil || errz != nil {
			fmt.Println("Error(s) converting coordinates: ", errx, erry, errz)
		}
		if errm != nil {
			fmt.Println("Error converting MEP value: ", errm)
		}

		out.Unpackr(j, float32(x), float32(y), float32(z), float32(m))
	}
	return out
}

//PtstoImgi converts points to an image
func PtstoImgi(pts *Matrixi, wid, hei int) *Matrixi {
	target := NewMatrixi(256, 256)
	for j := 0; j < pts.Rows; j++ {
		x := pts.Data[j][1]
		y := pts.Data[j][2]
		target.Data[y][x] = pts.Data[j][3]
	}
	return target
}

//PtstoImgf converts points to an image
func PtstoImgf(pts *Matrixf, wid, hei int) *Matrixf {
	target := NewMatrixf(256, 256)
	for j := 0; j < pts.Rows; j++ {
		x := int(Roundf(pts.Data[j][0]))
		y := int(Roundf(pts.Data[j][1]))
		target.Data[y][x] = pts.Data[j][3]
	}
	return target
}

//PtstoImgff converts points to an image
func PtstoImgff(pts *Matrixff, wid, hei int) *Matrixff {
	target := NewMatrixff(256, 256)
	for j := 0; j < pts.Rows; j++ {
		x := int(Roundff(pts.Data[j][0]))
		y := int(Roundff(pts.Data[j][1]))
		target.Data[y][x] = pts.Data[j][3]
	}
	return target
}
