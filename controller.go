package main

import (
	"context"
	"fmt"
	"goweb/framework"
	"net/http"
	"time"
)

func FooControllerHandler(c *framework.Context) error {

	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancle := context.WithTimeout(c.BaseContext(),
		time.Duration(1*time.Second))
	defer cancle()

	go func() {
		// 距离的业务逻辑
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)
		c.Json(http.StatusOK, "ok")

		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		// 出现异常
		fmt.Printf("panic %v\n", p)
	case <-finish:
		// 业务逻辑执行完成
		fmt.Printf("finish\n")
	case <-durationCtx.Done():
		// 超时时间到
		fmt.Printf("timeout\n")
	}

	return nil
}
