package main

import (
	"fmt"
	"goweb/framework"
	"net/http"
)

func FooControllerHandler(c *framework.Context) error {

	c.Json(http.StatusOK, "FooControllerHandler")
	return nil
	//finish := make(chan struct{}, 1)
	//panicChan := make(chan interface{}, 1)
	//
	//durationCtx, cancel := context.WithTimeout(c.BaseContext(),
	//	time.Duration(1*time.Second))
	//defer cancel()
	//
	//go func() {
	//	// 具体的义务逻辑
	//	defer func() {
	//		if p := recover(); p != nil {
	//			panicChan <- p
	//		}
	//	}()
	//
	//	time.Sleep(10 * time.Second)
	//	c.Json(http.StatusOK, "ok")
	//
	//	finish <- struct{}{}
	//}()
	//
	//select {
	//case p := <-panicChan:
	//	// 出现异常
	//	fmt.Printf("panic %v\n", p)
	//case <-finish:
	//	// 业务逻辑执行完成
	//	fmt.Printf("finish\n")
	//case <-durationCtx.Done():
	//	c.WriterMux().Lock()
	//	defer c.WriterMux().Unlock()
	//	// 超时时间到
	//	fmt.Printf("timeout\n")
	//	c.SetHasTimeout()
	//}
	//
	//return nil
}

func SubjectDelController(c *framework.Context) error {
	id := c.QueryInt("id", 0)
	c.Json(http.StatusOK, id)
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.Json(http.StatusOK, "ok")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	id := c.FormInt("id", -1)
	values := c.GetRequest().URL.Query()
	params := values["id"]
	fmt.Printf("params = %v\n", params)
	c.Json(http.StatusOK, id)
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.Json(http.StatusOK, "SubjectListController")
	return nil
}
