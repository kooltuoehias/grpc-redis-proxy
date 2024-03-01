package server

import (
	"context"
	"fmt"
	"time"

	pb "github.com/kooltuoehias/grpc-redis-client/grpc"
	"github.com/redis/go-redis/v9"
)

type RedisProxyGrpcServer struct {
	pb.UnimplementedRedisProxyServer
	redisClient *redis.Client
}

func NewServer(client *redis.Client) RedisProxyGrpcServer {
	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	_, err = client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client ready to Get and Set")

	return RedisProxyGrpcServer{redisClient: client}
}

func (server RedisProxyGrpcServer) Get(ctx context.Context, request *pb.RedisRequest) (*pb.RedisReply, error) {
	val, err := server.redisClient.Get(ctx, request.GetKey()).Result()
	if err != nil {
		return &pb.RedisReply{Value: err.Error(), Status: "failure"}, err
	}
	fmt.Println(request.Key, val)
	return &pb.RedisReply{Value: val, Status: "success"}, nil
}

func (server RedisProxyGrpcServer) Set(ctx context.Context, request *pb.RedisSetRequest) (*pb.RedisReply, error) {
	val, err := server.redisClient.Set(ctx, request.GetKey(),
		request.GetValue(),
		time.Duration(request.GetDurationInSecond())*time.Second).Result()
	if err != nil {
		return &pb.RedisReply{Value: err.Error(), Status: "failure"}, err
	}
	fmt.Println(request.Key, val, request.GetDurationInSecond())
	return &pb.RedisReply{Value: val, Status: "success"}, nil
}
