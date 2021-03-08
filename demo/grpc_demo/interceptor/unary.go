package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"runtime/debug"
)

func Filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("filter:  ", info)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v\n, stack: %v\n", r, debug.Stack())
		}
	}()
	return handler(ctx, req)
}

func LogReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Printf("request info: %v\n", req)
	log.Printf("server info: %s -> %s\n", info.Server, info.FullMethod)

	return handler(ctx, req)
}
