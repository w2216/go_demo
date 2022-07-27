package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	"time"
)

func main() {
	// 1. 用 rpc 连接服务器 -- Dial
	conn, err1 := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer conn.Close()

	// 2. 调用远程函数
	var reply string
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	err2 := conn.Call("hello.SayHello", "张三 "+timeStr, &reply)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(reply)
}

// jsonrpc总结
// 1. conn, err1 := rpc.Dial("tcp", "127.0.0.1:8080") 改成
// 1. conn, err1 := jsonrpc.Dial("tcp", "127.0.0.1:8080")

// 2. err2 := conn.Call("hello.SayHello", "张三 " + timeStr, &reply)
