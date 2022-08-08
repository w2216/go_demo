package main

import (
	"client/services/hello"
	"context"
	"fmt"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/go-micro/plugins/v3/registry/etcd"
)

func main() {
	// etcd
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	// 创建服务
	service := micro.NewService(
		micro.Registry(etcdRegistry),
	)

	// 初始化
	service.Init()

	// new 一个hello 服务
	clientHello := hello.NewSayService("hello", service.Client())

	// 发送请求
	res, err := clientHello.Hello(context.Background(), &hello.Request{
		Name: "张三...",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
