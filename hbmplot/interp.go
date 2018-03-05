package hbmplot

/*func lininterpi(Pt hbmutil.Point2i, Corners []hbmutil.Point2i) int {
	if len(Corners) != 4 {
		panic("Must input 4 corners for linear interpolation")
	}
	var x11 hbmutil.Point2i
	var x12 hbmutil.Point2i
	var x21 hbmutil.Point2i
	var x22 hbmutil.Point2i
	for _, v := range Corners {
		if v.X > Pt.X && v.Y > Pt.Y {
			x22 = *v
		}
		if v.X > Pt.X && v.Y < Pt.Y {
			x21 = *v
		}
		if v.X < Pt.X && v.Y > Pt.Y {
			x12 = *v
		}
		if v.X < Pt.X && v.Y < Pt.Y {
			x11 = *v
		}
	}
	interpmat := hbmutil.NewMatrixi(4, 4)
	interpmat.Unpackr(0, 1, x11.X, x11.Y, x11.X*x11.Y)
	interpmat.Unpackr(0, 1, x12.X, x12.Y, x12.X*x12.Y)
	interpmat.Unpackr(0, 1, x21.X, x21.Y, x21.X*x21.Y)
	interpmat.Unpackr(0, 1, x22.X, x22.Y, x22.X*x22.Y)
	coeffs := interpmat.Cramer(x11.V, x12.V, x21.V, x22.V)

	return coeffs[0] + Pt.X*coeffs[1] + Pt.Y*coeffs[2] + Pt.X*Pt.Y*coeffs[3]
}

func Squash(pts []hbmutil.Point3f, dim int) []hbmutil.Point2f {
	out := make([]hbmutil.Point2f, len(pts))
	switch dim {
	case 1:
		for i, v := range pts {
			out[i] = hbmutil.Point2f{v.Y, v.Z}
		}
	case 2:
		for i, v := range pts {
			out[i] = hbmutil.Point2f{v.X, v.Z}
		}
	case 3:
		for i, v := range pts {
			out[i] = hbmutil.Point2f{v.X, v.Y}
		}
	case -1:
		for i, v := range pts {
			out[i] = hbmutil.Point2f{v.Z, v.Y}
		}
	case -2:
		for i, v := range pts {
			out[i] = hbmutil.Point2f{v.Z, v.X}
		}
	case -3:
		for i, v := range pts {
			out[i] = hbmutil.Point2f{v.Y, v.X}
		}
	}
	return out
}

func BilinearInterp([]hbmutil.Point2f) *hbmutil.Matrixf {

}
*/
