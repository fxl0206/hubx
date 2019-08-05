export GOPATH=/root/apps/hubx/hubxs
go build -ldflags "-w -s" -o hubxs $GOPATH/src/aif.io/hubx/cmd/main.go
