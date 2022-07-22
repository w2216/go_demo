package controller

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// 流水线模型2
func Chan2Controller(ctx *gin.Context) {
	arrInt := []int{}
	for i := 1; i < 10; i++ {
		arrInt = append(arrInt, i)
	}

	outOne := producer2(arrInt...)
	outTwo1 := square2(outOne)
	outTwo2 := square2(outOne)
	outTwo3 := square2(outOne)

	result := []int{}
	for v := range merge(outTwo1, outTwo2, outTwo3) {
		result = append(result, v)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": result,
	})
}

// 放入chan
func producer2(num ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range num {
			out <- v
		}
	}()
	return out
}

// 计算值, 并返回chan
func square2(inCh chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range inCh {
			out <- v * v
			time.Sleep(time.Second)
		}
	}()
	return out
}

// 合并chan
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}
	wg.Add(len(cs))
	// FAN-IN
	for _, c := range cs {
		go collect(c)
	}
	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
