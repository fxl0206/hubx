set GOPATH=D:\gopaths\protoc

protoc --go_out=../../../../ *.proto
protoc --gofast_out=../../../../ *.proto
protoc --gogofast_out=../../../../../../ *.proto
D:\gopaths\protoc\bin\protoc --gogofast_out=../../../../../../ *.proto
git config --global core.autocrlf false