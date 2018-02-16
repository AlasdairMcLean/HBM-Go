package hbmutil

//Pointi defines an interface for 2d and 3d int points
type Pointi interface {
}

//Pointf defines an interface for 2d and 3d float32 points
type Pointf interface {
}

//Pointff defines an interface for 2d and 3d float64 points
type Pointff interface {
}

//Point3i defines a 3d integer point
type Point3i struct {
	X, Y, Z, V int
}

//Point3f defines a 3d float32 point
type Point3f struct {
	X, Y, Z, V float32
}

//Point3ff defines a 3d float64 point
type Point3ff struct {
	X, Y, Z, V float64
}

//Point2i defines a 2d integer point
type Point2i struct {
	X, Y, V int
}

//Point2f defines a 2d float32 point
type Point2f struct {
	X, Y, V float32
}

//Point2ff defines a 2d float64 point
type Point2ff struct {
	X, Y, V float64
}

type PointCloud2f struct {
	Pts    []Point2f
	center float32
}

//func (PointCloud2f)
