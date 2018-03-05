package hbmcluster

import (
	"fmt"
	"time"

	"hbm/util"
)

//ClusterFromFile returns an array of nodes based off of a csv file which have been clustered based on either default or inputted values
//Recommended default values: numrep=200; NodeThresh=4; NumNodes=50
func ClusterFromFile(fulldatafilepat string, MEPcol, NumNodes, numrep int, NodeThresh float32) []*ClusterNode {

	datamat, err := hbmutil.ReadMEP(fulldatafilepat, 0, 1, 2, 3)
	if err != nil {
		panic(err)
	}
	data := datamat.Data

	nodecount := 0 // first node is located at the first point
	AllNodes := make([]*ClusterNode, 0, NumNodes)
	AllNodes = append(AllNodes, NewNodePts(data[0][0], data[0][1], data[0][2], data[0][3], numrep))
	pointsproc := 0                  // counter for total points processed
	TESTT1 := time.Now()             // This is the "tic function" from matlab
	for n := 1; n < len(data); n++ { // starting off from the second point
		nx := data[n][0] // offload the x,y,z values for ease later
		ny := data[n][1]
		nz := data[n][2]
		na := data[n][3]
		nodefound := false //while loop condition- iterate until we find or create a node for every point
		nodenum := 0       //iterator variable for the while loop
		for nodefound == false {
			noi := AllNodes[nodenum]                           //make a variable for the node for simplicity
			if noi.ToPoint(nx, ny, nz) < float32(NodeThresh) { //if the distance to the point from each node is less than the threshold
				noi.Add(nx, ny, nz, na)
				nodefound = true
				pointsproc = pointsproc + 1
			}
			if nodenum == nodecount && nodefound == false {
				nodecount = nodecount + 1
				AllNodes = append(AllNodes, NewNodePts(nx, ny, nz, na, numrep))
				nodefound = true
				pointsproc = pointsproc + 1
			}
			nodenum = nodenum + 1
		}
	}

	fmt.Println("Time Taken: ", time.Since(TESTT1))
	// This is lines 218 onwards

	//[p,t,~]:= elipfit(nodepts(:,1:2)){ // not sure if elipfit is a thing in go
	//xq := hbmutil.Linspacef((hbmutil.min(p)), hbmutil.max(p), 256)
	//yq := hbmutil.Linspacef(hbmutil.min(t), hbmutil.max(t), 256) // is linspace a thing in go ask ally
	//[Xq,Yq]=meshgrid(xq,yx); // Creates a meshgrid
	//Vq=griddata(p,t,nodepts)(:,5),Xq,Yq,'cubic') // idk how to make that a string do i just say string right afterwards like above?
	//allpts := hbmutil.NewMatrixf(6, 103*104) // topographical idea of how this looks like
	//allptct := 0
	//cmall:= jet(round(max(max(Vq(:,1)))))

	/*
	   for i:= 1:len(Vq(:,1)) {
	   	for j:= 1:len(Vq(1,:)) {
	   		if  isnan(Vq(i,j)) {
	   			Vq(i,j):= 0;
	   		}
	   		if Vq(i,j) < 50{
	   			Vq(i,j):=0;
	   		}
	   		allptct:=allptct+1;
	           allpts(:,allptct)=[i,j,Vq(i,j),cmall(round(Vq(i,j))+1,:)];
	   	}
	   }*/
	return AllNodes
}
