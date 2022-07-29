package main

import (
	"context"
	"fmt"

	"client/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 创建grpc连接
	grpcClient, err1 := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer grpcClient.Close()

	// 2. 创建一个hello连接
	clientHello := services.NewHelloClient(grpcClient)

	// 3. 发送请求
	req := &services.HelloReq{
		Name: "张三",
	}
	res := &services.HelloRes{}
	res, err2 := clientHello.SayHello(context.Background(), req)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Printf("%#v\n", res.Msg)
}
