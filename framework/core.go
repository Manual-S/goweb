package framework

import (
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{
		router: map[string]ControllerHandler{},
	}
}

// ServeHTTP 自定义的ServeHTTP
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 处理路由
	ctx := NewContext(r, w)
	handler := c.router["foo"]

	handler(ctx)
}

func (c *Core) Get(path string, handler ControllerHandler) {
	c.router[path] = handler
}
