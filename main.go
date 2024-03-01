package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/kooltuoehias/grpc-redis-client/grpc"
	"github.com/kooltuoehias/grpc-redis-client/redis"
	"github.com/kooltuoehias/grpc-redis-client/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRedisProxyServer(grpcServer, server.NewServer(redis.RedisConnection()))
	reflection.Register(grpcServer)
	fmt.Printf("localhost:%d \n", *port)
	grpcServer.Serve(lis)
}
