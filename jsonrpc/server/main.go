package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Hello struct {
	Name string
}

// 服务方法
func (h Hello) SayHello(req string, res *string) error {
	fmt.Println(req)
	*res = "你好" + req
	return nil
}

// server端
func main() {
	// 1. 注册一个服务
	err1 := rpc.RegisterName("hello", &Hello{})
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	// 2. 设置监听
	listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer listener.Close()

	for {
		fmt.Println("server: 开始工作...")
		// 3. 接收连接
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
			return
		}

		// 4. 绑定业务
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}
}

// jsonrpc总结
// 1. err1 := rpc.RegisterName("hello", &Hello{})
// 2. listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
// 3. conn, err3 := listener.Accept()

// 4. go rpc.ServeConn(conn) 改成
// 4. go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
