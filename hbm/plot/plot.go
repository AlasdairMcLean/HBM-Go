//Package hbmplot establishes plotting functions for use in gtk/gdk/cairo using gotk3/gotk3 bindings
package hbmplot

import (
	"fmt"
	"math"

	"hbm/util"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

//Drawim renders an image into a given DrawingArea.
func Drawim(da *gtk.DrawingArea, img *hbmutil.Matrixff, cmap string) {
	var cmat *hbmutil.Matrixf
	fmt.Println(cmap)
	switch cmap {
	case "gsc":
		cmat = hbmutil.Gscmap(int(math.Floor(img.Maxa())+1), 1)
	case "hot":
		cmat = hbmutil.Hotmap(int(math.Floor(img.Maxa())+1), 1)
	default:
		pmsg := fmt.Sprintf("%v is not a valid colormap. Try 'gsc' for greyscale or 'hot' for an even-luminated map.", cmap)
		panic(pmsg)
	}
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) { //draw each pixel specified by following fcn:
		for j := 0; j < img.Rows; j++ { //for each row
			for i := 0; i < img.Cols; i++ { //for each col
				cr.Rectangle(float64(i), float64(j), 5, 5) //connect borders of a rectangle (pixel). Currently 5x5px for visibility, but will go back to 1x1 once interpolation is implemented
				val := img.Data[j][i]
				r := float64(cmat.Data[int(val)][0])
				g := float64(cmat.Data[int(val)][1])
				b := float64(cmat.Data[int(val)][2])
				cr.SetSourceRGB(r, g, b) //set the rgb in greyscale based on the intensity value
				//fill in the pixel
				cr.Fill()
			}
		}
	})
}

//ExpandCanvas expands the size of the canvas in the screen;
func ExpandCanvas(da *gtk.DrawingArea, img *hbmutil.Matrixff, scale int) {
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetOperator(cairo.OPERATOR_CLEAR)
		cr.Paint()
		cr.SetOperator(cairo.OPERATOR_SOURCE)
		for j := 0; j < img.Rows*scale; j += scale {
			for i := 0; i < img.Cols*scale; i += scale {
				cr.Rectangle(float64(i), float64(j), float64(scale), float64(scale))
				cr.SetSourceRGB(float64(img.Data[j/scale][i/scale]), float64(img.Data[j/scale][i/scale]), float64(img.Data[j/scale][i/scale]))
				cr.Fill()
			}
		}
	})
}

//ZoomCanvas Zooms in on the size of the canvas in the screen without changing the size of the canvas
func ZoomCanvas(da *gtk.DrawingArea, img *hbmutil.Matrixff, canvsc int, cmap string) {
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		var cmat *hbmutil.Matrixf
		fmt.Println(cmap)
		switch cmap {
		case "gsc":
			cmat = hbmutil.Gscmap(int(math.Floor(img.Maxa())+1), 1)
		case "hot":
			cmat = hbmutil.Hotmap(int(math.Floor(img.Maxa())+1), 1)
		default:
			pmsg := fmt.Sprintf("%v is not a valid colormap. Try 'gsc' for greyscale or 'hot' for an even-luminated map.", cmap)
			panic(pmsg)
		}
		scale := math.Pow(2, float64(canvsc-1))
		fmt.Println(scale)
		fmt.Println(img.Rows, img.Cols)
		var (
			rowstart, rowend, colstart, colend int
		)
		if canvsc > 1 {
			ps, pe := gethalf(scale)
			rowstart = int(float64(img.Rows) * ps)
			rowend = int(float64(img.Rows) * pe)
			colstart = int(float64(img.Cols) * ps)
			colend = int(float64(img.Cols) * pe)
			fmt.Println("scale ", scale, "; powstart ", pe, "; powend ", ps)
		} else {
			rowstart = 0
			rowend = img.Rows
			colstart = 0
			colend = img.Cols
		}
		fmt.Println(rowstart, rowend, colstart, colend)
		for j, jc := rowstart, 0; j < rowend; j++ {
			for i, ic := colstart, 0; i < colend; i++ {
				cr.Rectangle(float64(ic), float64(jc), 256/float64((colend-colstart)), 256/float64((rowend-rowstart)))
				val := img.Data[j][i]
				r := float64(cmat.Data[int(val)][0])
				g := float64(cmat.Data[int(val)][1])
				b := float64(cmat.Data[int(val)][2])
				cr.SetSourceRGB(r, g, b) //set the rgb in greyscale based on the intensity value
				cr.Fill()

				ic += 256 / (colend - colstart)
			}
			jc += 256 / (rowend - rowstart)
		}
	})
}

func gethalf(scale float64) (float64, float64) {
	var (
		s float64
		e float64 = 1
	)
	for i := 1; i < int(scale); i++ {
		fmt.Println(s, e)
		s = 0.5 - 0.5*(0.5-s)
		e = 0.5*(e-0.5) + 0.5
	}
	return s, e
}

/*
func Zoomin(da *gtkDrawingArea, img *hbmutil.Matrixff, scale int) {
	nativeWid:=img.Cols
	nativeHei:=img.Rows

	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		for i := 0; i < img.Cols; i += scale {
			for j := 0; j < img.Rows; j += scale {
				cr.Rectangle(float64)
			}
		}
	})
}*/
