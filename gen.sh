export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
protoc --go_out=.  --go-grpc_out=.  ./proto/service.proto