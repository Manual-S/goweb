// Package framework 批量前缀实现
package framework

type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

type Group struct {
	core   *Core
	prefix string
}

func NewGroup(c *Core, p string) IGroup {
	return &Group{
		core:   c,
		prefix: p,
	}
}

func (g *Group) Get(uri string, handler ControllerHandler) {
	g.core.Get(g.prefix+uri, handler)
}
func (g *Group) Post(uri string, handler ControllerHandler) {
	g.core.Post(g.prefix+uri, handler)
}
func (g *Group) Put(uri string, handler ControllerHandler) {
	g.core.Put(g.prefix+uri, handler)
}
func (g *Group) Delete(uri string, handler ControllerHandler) {
	g.core.Delete(g.prefix+uri, handler)
}
