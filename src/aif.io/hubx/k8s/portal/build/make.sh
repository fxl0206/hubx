export GOPATH=/root/apps/hubx/hubxs
go build -ldflags "-w -s" -o hubx $GOPATH/src/aif.io/hubx/k8s/portal/cmd/main.go
