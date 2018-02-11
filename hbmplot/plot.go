//Package hbmplot establishes plotting functions for use in gtk/gdk/cairo using gotk3/gotk3 bindings
package hbmplot

import (
	"../hbmutil"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

//Drawim renders an image into a given DrawingArea.
func Drawim(da *gtk.DrawingArea, img *hbmutil.Matrixff) {
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		for j := 0; j < img.Rows; j++ {
			for i := 0; i < img.Cols; i++ {
				cr.SetSourceRGB(float64(img.Data[j][i]), float64(img.Data[j][i]), float64(img.Data[j][i]))
				cr.Rectangle(float64(i), float64(j), 5, 5)
				cr.Fill()
			}
		}
	})

}
