package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"gotools/demo/grpc_demo/auth"
	"gotools/demo/grpc_demo/interceptor"
	hello "gotools/demo/grpc_demo/proto"
	pb "gotools/demo/grpc_demo/proto"
	"io/ioutil"
	"log"
	"net"
)

var (
	Address = ":8080"
)

type HelloService struct {
	auth *auth.Authentication
}

func NewHelloService() *HelloService {
	return &HelloService{
		auth: &auth.Authentication{
			Username: "hello",
			Password: "world",
		},
	}
}

func (h HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	if err := h.auth.Auth(ctx); err != nil {
		return nil, err
	}
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

func main() {

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %s\n", err.Error())
	}

	crt_file := `H:\zen0fpy\gotools\demo\grpc_demo\keys\server.pem`
	key_file := `H:\zen0fpy\gotools\demo\grpc_demo\keys\server.key`
	ca_file := `H:\zen0fpy\gotools\demo\grpc_demo\keys\ca.pem`

	certificate, err := tls.LoadX509KeyPair(crt_file, key_file)
	if err != nil {
		log.Fatalf("Failed to load cert and key file, %s\n", err.Error())
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca_file)
	if err != nil {
		log.Fatal(err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append pem")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	s := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(interceptor.LogReq, interceptor.Filter),
	))

	helloService := NewHelloService()
	hello.RegisterHelloServer(s, helloService)

	log.Printf("Listen on " + Address)

	s.Serve(listen)
}
