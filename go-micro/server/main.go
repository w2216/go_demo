package main

import (
	"context"
	"fmt"
	"server/services/hello"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/go-micro/plugins/v3/registry/etcd"
)

type sayHandler struct {
	hello.SayHandler
}

func (h *sayHandler) Hello(ctx context.Context, req *hello.Request, res *hello.Response) error {
	fmt.Println(req)
	res.Msg = "你好" + req.GetName()
	return nil
}

func main() {
	// etcd
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	// 创建服务 hello
	service := micro.NewService(
		// 服务名称
		micro.Name("hello"),
		micro.Version("latest"),
		micro.Registry(etcdRegistry),
	)

	// 初始化
	service.Init()

	// 注册 hello handler
	hello.RegisterSayHandler(service.Server(), new(sayHandler))

	// 开启服务
	service.Run()
}
