package hbmcluster

import (
	"math"

	"hbm/util"
)

//ClusterNode is a self-updating struct that keeps track of the points belonging to it as well as the centroid.
type ClusterNode struct {
	X, Y, Z float32
	Pts     hbmutil.Matrixf
	Curpts  int
	Maxpts  int
}

//MakeNode is the constructor function for making a new empty clusternode
func MakeNode(cap int) *ClusterNode {
	out := ClusterNode{Pts: *hbmutil.NewMatrixf(cap, 4), Maxpts: cap, Curpts: 0}
	return &out
}

//NewNodePts is a constructor function where the first point is initialized
func NewNodePts(x, y, z, v float32, cap int) *ClusterNode {
	out := ClusterNode{Pts: *hbmutil.NewMatrixf(cap, 4), Maxpts: cap, Curpts: 1}
	out.Pts.Unpackr(0, x, y, z, v)
	return &out
}

//ToOrigin will quickly calculate the distance between the centroid of the node to the origin (0,0,0)
func (cn *ClusterNode) ToOrigin() float32 {
	return float32(math.Sqrt(math.Pow(float64(cn.X), 2) + math.Pow(float64(cn.Y), 2) + math.Pow(float64(cn.Z), 2)))
}

//ToPoint will return the distance between the given point and the centroid of the node
func (cn *ClusterNode) ToPoint(x, y, z float32) float32 {
	return float32(math.Sqrt(math.Pow(float64(cn.X-x), 2) + math.Pow(float64(cn.Y-y), 2) + math.Pow(float64(cn.Z-z), 2)))
}

//Add is a constructor function for adding a point to the node, which triggers the recalculation of the centroid
func (cn *ClusterNode) Add(x, y, z, amp float32) {
	cn.Pts.Data[cn.Curpts][0] = x
	cn.Pts.Data[cn.Curpts][1] = y
	cn.Pts.Data[cn.Curpts][2] = z
	cn.Pts.Data[cn.Curpts][3] = amp
	cn.Curpts++
	cn.X = hbmutil.Avgfl(cn.Pts.Getcol(0)[:cn.Curpts])
	cn.Y = hbmutil.Avgfl(cn.Pts.Getcol(1)[:cn.Curpts])
	cn.Z = hbmutil.Avgfl(cn.Pts.Getcol(2)[:cn.Curpts])
}
