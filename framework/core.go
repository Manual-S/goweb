package framework

import (
	"fmt"
	"net/http"
)

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}

// ServeHTTP 自定义的ServeHTTP
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 简单的打印
	fmt.Printf("hello word")
}
