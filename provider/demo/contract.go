// Package demo 服务接口文件
package demo

// Key 注意这里的Key首字母要大写 表示可导出的
const Key = "web:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
