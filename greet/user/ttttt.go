package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const url = "http://localhost:8888/user/add"
const contentType = "application/json"

var wg = sync.WaitGroup{}

func main() {

	count := 1000 // 最大支持并发
	sum := 10000  // 任务总数

	c := make(chan struct{}, count) // 控制任务并发的chan
	defer close(c)

	for i := 0; i < sum; i++ {
		var s = "{\"id\":0,\"name\":\"" + "张三" + strconv.Itoa(100000+i) +
			"\",\"phone\":\"" + strconv.Itoa(18100000000+i) + "\",\"password\":\"123456\"}"
		wg.Add(1)
		c <- struct{}{} // 作用类似于waitgroup.Add(1)
		go func(j int) {
			go Run(s, c)
		}(i)
	}
	wg.Wait()
}

func Run(s string, c chan struct{}) {
	defer wg.Done()
	fmt.Println(s)
	_, _ = http.Post(url, contentType, strings.NewReader(s))
	time.Sleep(time.Second * 2)
	<-c // 执行完毕，释放资源
}
