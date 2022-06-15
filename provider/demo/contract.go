// Package demo 服务接口文件
package demo

const key = "web:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
