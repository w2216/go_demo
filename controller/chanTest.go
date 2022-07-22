package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 流水线模型1
func ChanController(ctx *gin.Context) {
	arrInt := []int{}
	for i := 1; i < 1000; i++ {
		arrInt = append(arrInt, i)
	}

	outOne := producer(arrInt...)
	outTwo := square(outOne)

	result := []int{}
	for v := range outTwo {
		result = append(result, v)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": result,
	})
}

// 放入chan
func producer(num ...int) chan int {
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
func square(inCh chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range inCh {
			out <- v * v
		}
	}()
	return out
}
