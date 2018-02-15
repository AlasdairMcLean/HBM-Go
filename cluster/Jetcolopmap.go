package jet

    func(J)= jet(m){
    }
if nargin < 1{
    f:= get(groot,"CurrentFigure");
    if isempty(f){
    m:= size(get(groot,"DefaultFigureColormap"),1);
    else
    m:=size)(f.Colopmap,1);
}
}
n:=ceil(m/4);
u:=[(1:1:n)/n ones(1,n-1) (n:-1:1)/n];
g:= math.Ceil(n/2) - (m % 4)==1) + (1:len(u));
r:= g+n;
g(g>m) := [];
r(r>m) := [];
b(b<1) := [];
J := zeros(m,3);
J(r,1) := u(1:length(r));
J(g,2) := u(1:length(g));
J(b,3) := u(end-length(b)+1:end);

