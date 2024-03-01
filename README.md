# grpc-redis-proxy
A vanilla grpc servier in Golang works as a redis proxy

## protoc in docker
```docker build . -t protoc-gen```    
```docker run --rm -it   --mount src="$(pwd)",target=/app,type=bind protoc-gen```    

## redis
A redis container shall be running on localhost:6379    

## run
go run main.go

## test
grpcurl  -plaintext  -d '{"key":"foo"}' localhost:50051  RedisProxy/get    
{
  "value": "bar",
  "status": "success"
}    

grpcurl  -plaintext  -d '{"key":"foo","value":"newfoo"}' localhost:50051  RedisProxy/set    
{
  "value": "OK",
  "status": "success"
}    

grpcurl  -plaintext  -d '{"key":"foo"}' localhost:50051  RedisProxy/get    
{
  "value": "newfoo",
  "status": "success"
}    

## doc
[Go | gRPC](https://grpc.io/docs/languages/go/)    
[grpcurl](https://github.com/fullstorydev/grpcurl)    
[Go | gRPC | examples](https://github.com/grpc/grpc-go/tree/master/examples)