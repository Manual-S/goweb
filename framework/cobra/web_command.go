package cobra

import (
	"github.com/robfig/cron"
	"goweb/framework"
	"log"
)

func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

func (c *Command) GetContainer() framework.Container {
	// 这里是返回Root中的container
	return c.Root().container
}

// CronSpec 用来保存Cron命令的信息 用于展示
type CronSpec struct {
	Type        string
	Cmd         *Command
	Spec        string
	ServiceName string
}

// AddCronCommand 创建一个定时任务 调用方法
// AddCronCommand("* * * * * *",demo.FooCommand)
func (c *Command) AddCronCommand(spec string, cmd *Command) {
	root := c.Root()
	if root.CronClient == nil {
		// 初始化一个CronClient
		root.CronClient = cron.New()
		root.CronSpecs = []CronSpec{}
	}

	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type: "normal-cron",
		Cmd:  cmd,
		Spec: spec,
	})

	var cronCmd Command
	ctx := root.Context()

	cronCmd = *cmd
	cronCmd.args = []string{}
	cronCmd.SetParentNull()
	cronCmd.HasParent()
	cronCmd.SetContainer(root.GetContainer())

	root.CronClient.AddFunc(spec, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("CronClient.AddFunc error %v\n", err)
			}
		}()

		err := cronCmd.ExecuteContext(ctx)

		if err != nil {
			log.Printf("cronCmd.ExecuteContext error %v\n", err)
		}
	})
}
