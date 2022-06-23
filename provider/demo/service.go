package demo

import "goweb/framework"

type DemoService struct {
	Service // 匿名结构体变量 表示一种实现关系

	c framework.Container
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)

	return &DemoService{c: c}, nil
}

// GetFoo 实现Service的接口
func (d *DemoService) GetFoo() Foo {
	return Foo{
		Name: "my name is lili",
	}
}
