package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	ctx            context.Context
	request        *http.Request
	responseWriter http.ResponseWriter
	writeMux       *sync.Mutex
	hasTimeout     bool // 是否超时标记位
	handlers       []ControllerHandler
	index          int // 表示执行到了那个函数
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

// request相关函数

func (c *Context) QueryInt(key string, def int) int {
	mapRes := c.QueryAll()
	if valus, ok := mapRes[key]; ok {
		if len(valus) > 0 {
			res, err := strconv.Atoi(valus[len(valus)-1])
			if err != nil {
				return def
			}

			return res
		}
	}

	return def
}

// QueryString 查询请求中的string参数
func (c *Context) QueryString(key string, def string) string {
	mapRes := c.QueryAll()
	if valus, ok := mapRes[key]; ok {
		if len(valus) > 0 {
			res := valus[len(valus)-1]
			return res
		}
	}

	return def
}

func (c *Context) QueryArray(key string, def []string) []string {
	return nil
}

func (c *Context) QueryAll() map[string][]string {
	if c.request != nil {
		return map[string][]string(c.request.URL.Query())
	}

	return map[string][]string{}
}

func (c *Context) FormInt(key string, def int) int {
	mapRes := c.FormAll()
	if valus, ok := mapRes[key]; ok {
		if len(valus) > 0 {
			res, err := strconv.Atoi(valus[len(valus)-1])
			if err != nil {
				return def
			}

			return res
		}
	}

	return def
}

func (c *Context) FormAll() map[string][]string {
	if c.request != nil {
		return map[string][]string(c.request.PostForm)
	}

	return map[string][]string{}
}

// response相关函数

// Json 返回json结构
func (c *Context) Json(status int, obj interface{}) error {
	if c.HasTimeout() {
		// 已经超时了
		return nil
	}
	c.responseWriter.Header().Set("Content-Type", "application/json")
	c.responseWriter.WriteHeader(status)

	data, err := json.Marshal(obj)
	if err != nil {
		c.responseWriter.WriteHeader(http.StatusInternalServerError)
		return err
	}

	c.responseWriter.Write(data)
	return nil
}

func (c *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (c *Context) Text(status int, obj interface{}) error {
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
