package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// Token认证
// 要实现对每个gRPC方法进行认证，需要实现grpc.PerRPCCredentials接口

type Authentication struct {
	Username string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"user":     a.Username,
		"password": a.Password,
	}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("get from incoming context failed")
	}
	var appid string
	var appKey string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}

	if val, ok := md["password"]; ok {
		appKey = val[0]
	}

	if appid != a.Username && appKey != a.Password {
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}
	return nil

}
