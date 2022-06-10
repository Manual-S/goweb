package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]*Tree
}

func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree() // gin框架采用的是slice 这里采用的是map结构
	router["POST"] = NewTree()
	router["DELETE"] = NewTree()
	router["PUT"] = NewTree()

	return &Core{
		router: router,
	}
}

// ServeHTTP 自定义的ServeHTTP
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("path = %v", r.URL.Path)

	// 处理路由
	ctx := NewContext(r, w)
	router := c.FindRouterByRequest(r)
	if router == nil {
		ctx.Json(http.StatusOK, "not find router")
		return
	}
	err := router(ctx)
	if err != nil {
		ctx.Json(http.StatusInternalServerError, "inner error")
		return
	}
}

func (c *Core) Get(path string, handler ControllerHandler) {
	err := c.router["GET"].AddRouter(path, handler)
	if err != nil {
		log.Fatalf("AddRouter error GET:%v", err)
	}
}

func (c *Core) Post(path string, handler ControllerHandler) {
	err := c.router["POST"].AddRouter(path, handler)
	if err != nil {
		log.Fatalf("AddRouter error POST:%v", err)
	}
}

func (c *Core) Delete(path string, handler ControllerHandler) {
	err := c.router["DELETE"].AddRouter(path, handler)
	if err != nil {
		log.Fatalf("AddRouter error Delete:%v", err)
	}
}

func (c *Core) Put(path string, handler ControllerHandler) {
	err := c.router["PUT"].AddRouter(path, handler)
	if err != nil {
		log.Fatalf("AddRouter error Put:%v", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouterByRequest(r *http.Request) ControllerHandler {
	path := r.URL.Path
	method := r.Method

	methodHandlers, ok := c.router[method]
	if !ok {
		// 找不到对应的路由
		log.Printf("FindRouterByRequest error not find router")
		return nil
	}

	return methodHandlers.FindHandler(path)
}
