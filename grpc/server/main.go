package main

import (
	"context"
	"fmt"
	"net"

	"server/services"

	"google.golang.org/grpc"
)

type Hello struct {
	services.HelloServer
}

func (Hello) SayHello(c context.Context, req *services.HelloReq) (*services.HelloRes, error) {
	return &services.HelloRes{
		Msg: "你好" + req.Name,
	}, nil
}

func main() {
	// 1. 创建server
	grpcServer := grpc.NewServer()

	// 2. 注册Hello服务
	services.RegisterHelloServer(grpcServer, new(Hello))

	// 3. 监听端口
	listener, err1 := net.Listen("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer listener.Close()

	// 4. 绑定服务
	grpcServer.Serve(listener)

}
