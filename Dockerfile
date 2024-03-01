FROM alpine:3 

RUN apk add --no-cache go protoc git protobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

WORKDIR /app
COPY . /app

CMD ["sh", "gen.sh"]

#LAZYDOG: docker run --rm --mount src="$(pwd)",target=/app,type=bind  grpc-redis-proxy
#LAZYDOG: docker run --rm -it   --mount src="$(pwd)",target=/app,type=bind  grpc-redis-proxy
#LAZYDOG: ./grpcurl  -plaintext  -d '{"key":"foo"}' localhost:50051  RedisProxy/get
#LAZYDOG: ./grpcurl  -plaintext  -d '{"key":"foo","value":"newfoo"}' localhost:50051  RedisProxy/set
