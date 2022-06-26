package cobra

import "goweb/framework"

func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

func (c *Command) GetContainer() framework.Container {
	// 这里是返回Root中的container
	return c.Root().container
}
