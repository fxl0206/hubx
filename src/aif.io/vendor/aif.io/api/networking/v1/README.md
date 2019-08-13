set GOPATH=D:\gopaths\protoc

protoc --go_out=../../../../ *.proto
protoc --gofast_out=../../../../ *.proto
protoc --gogofast_out=. *.proto