package cluster

// Converted Clustering Code
func(vq)= BasicCluster(fulldatafilepat string , hMEPcol int , arguments ...int){

    switch len(arguments) {
	case 0:
		numrep:= 200; NODE_THRESH:= 4; NUM_NODES:=50 // and and is "they both need to be true or else theyre false"
		
	case 1:
		numrep:= 200; NODE_THRESH:= varargin[0]; NUM_NODES:=50
		
	case 2:
		numrep:= 200; NODE_THRESH:= varargin[0]; NUM_NODES:=varargin[1]
		
	case 3:
		numrep:= varargin[2]; NODE_THRESH:= varargin[0]; NUM_NODES:=varargin[1]
	}
	 default:
		 panic ("Too many argumrnts: Use the full path of the data, the numeric column the MEP values are in, and optionally the distance threshold between nodes, the number of nodes, and finally the number of repetitions at each node.")
}
oldnewcompplots:=false;
nodecount:=1;
AllNodes:= NodeArray(ClusterNode(datafile(1,1),datafile(1,2),datafile(1,3),datafile(1,MEPcol),numrep),NUM_NODES)
pointsproc:=0; // counter for total points processed
t1:=time.Now()// This is the "tic function" from matlab
telapsed:=time.Since(t1)
for n:=2:length(datafile(:,1)){// starting off from the second point
nx:=datafile(n,1); // offload the x,y,z values for ease later
ny:=datafile(n,2);
nz:=datafile(n,3);
na= datafile(n,MEPcol);
nodefound:= false; //while loop condition- iterate until we find or create a node for every point
nodenum:=1; //iterator variable for the while loop
for nodefound==false {
	noi:=(AllNodes(nodenum).values)//make a variable for the node for simplicity
	if noi.toPoint(nx,ny,nz) < NODE_THRESH { //if the distance to the point from each node is less than the threshold
		AllNode(nodenum).values=noi.nodeadd(nx,ny,nz,na)
		nodefound=true; pointsproc= pointsproc+1 
	}
}
if nodenum==nodecount && nodefound== false: }
nodecount= nodecount+1 
AllNodes(nodecount)+1. values= ClusterNode(nx,ny,nz,na,numrep);
nodefound=true; pointsproc = pointsproc+1
}
nodenum=nodenum+1;{
}

// This is lines 218 onwards
[p,t,~]:= elipfit(nodepts(:,1:2)){ // not sure if elipfit is a thing in go 
xq:=linspace((min(p)),max(p),256); yq:=linspace(min(t),max(t),256)// is linspace a thing in go ask ally
[Xq,Yq]=meshgrid(xq,yx); // Creates a meshgrid 
Vq=griddata(p,t,nodepts)(:,5),Xq,Yq,'cubic') // idk how to make that a string do i just say string right afterwards like above? 
allpts:= zeros(6,103*104); // topographical idea of how this looks like 
allptct:=0
cmall:= jet(round(max(max(Vq(:,1)))))
}
for i:= 1:length(Vq(:,1)) {
for j:= 1:length(Vq(1,:))
if  isnan(Vq(i,j))
Vq(i,j):= 0;
}
if Vq(i,j) < 50{
	Vq(i,j):=0;
}
allptct:=allptct+1;{
           allpts(:,allptct)=[i,j,Vq(i,j),cmall(round(Vq(i,j))+1,:)];
}
