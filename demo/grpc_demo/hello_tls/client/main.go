package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gotools/demo/grpc_demo/auth"
	hello "gotools/demo/grpc_demo/proto"
	"io/ioutil"
	"log"
)

const (
	Address = "127.0.0.1:8080"
)

func main() {

	crt_file := `H:\zen0fpy\gotools\demo\grpc_demo\keys\client.pem`
	key_file := `H:\zen0fpy\gotools\demo\grpc_demo\keys\client.key`
	ca_file := `H:\zen0fpy\gotools\demo\grpc_demo\keys\ca.pem`

	auth := auth.Authentication{
		Username: "hello",
		Password: "world",
	}

	certificate, err := tls.LoadX509KeyPair(crt_file, key_file)
	if err != nil {
		log.Fatalf("Load key and crt file filed, %s\n", err.Error())
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca_file)
	if err != nil {
		log.Fatalf("Failed to load ca.crt, %s\n", err.Error())
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs.")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("failed to connect: %s\n", err)
	}
	defer conn.Close()

	cli := hello.NewHelloClient(conn)
	resp, err := cli.SayHello(context.Background(), &hello.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("call SayHello failed, %s\n", err.Error())
	}

	fmt.Printf("result: %s\n", resp.GetMessage())

}
