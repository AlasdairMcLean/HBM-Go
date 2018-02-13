//Package hbmplot establishes plotting functions for use in gtk/gdk/cairo using gotk3/gotk3 bindings
package hbmplot

import (
	"../hbmutil"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

//Drawim renders an image into a given DrawingArea.
func Drawim(da *gtk.DrawingArea, img *hbmutil.Matrixff) {
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) { //draw each pixel specified by following fcn:
		for j := 0; j < img.Rows; j++ { //for each row
			for i := 0; i < img.Cols; i++ { //for each col
				cr.Rectangle(float64(i), float64(j), 5, 5)                                                 //connect borders of a rectangle (pixel). Currently 5x5px for visibility, but will go back to 1x1 once interpolation is implemented
				cr.SetSourceRGB(float64(img.Data[j][i]), float64(img.Data[j][i]), float64(img.Data[j][i])) //set the rgb in greyscale based on the intensity value
				//fill in the pixel
				cr.Fill()
			}
		}
	})

}
