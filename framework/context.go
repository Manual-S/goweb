package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	IRequest
	IResponse

	ctx            context.Context
	request        *http.Request
	responseWriter http.ResponseWriter
	writeMux       *sync.Mutex
	hasTimeout     bool // 是否超时标记位
	handlers       []ControllerHandler
	index          int // 表示执行到了那个函数
	params         map[string]string
}

// NewContext 构造函数
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		writeMux:       &sync.Mutex{},
		index:          -1,
	}
}

// base功能

func (c *Context) WriterMux() *sync.Mutex {
	return c.writeMux
}

func (c *Context) GetRequest() *http.Request {
	return c.request
}

func (c *Context) GetResponse() http.ResponseWriter {
	return c.responseWriter
}

// SetHasTimeout 设置context的超时时间
func (c *Context) SetHasTimeout() {
	c.hasTimeout = true
}

// HasTimeout 查看一个context的超时时间
func (c *Context) HasTimeout() bool {
	return false
}

// Context要实现标准context的接口

func (c *Context) BaseContext() context.Context {
	return c.request.Context()
}

func (c *Context) Done() <-chan struct{} {
	return nil
}

func (c *Context) Err() error {
	return nil
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (c *Context) Value(key interface{}) interface{} {
	return nil
}

// Next 实现中间件的链路调用
func (c *Context) Next() error {
	c.index++
	if c.index < len(c.handlers) {
		err := c.handlers[c.index](c)
		if err != nil {
			return err
		}
		// 注意 这里千万不要写c.index++
	}
	return nil
}

func (c *Context) SetHandlers(handlers []ControllerHandler) {
	c.handlers = append(c.handlers, handlers...)
}

func (c *Context) SetParams(hash map[string]string) {
	c.params = hash
}
