package cobra

import (
	"github.com/robfig/cron/v3"
	"goweb/framework/contract"
	"log"
	"time"
)

func (c *Command) AddDistributedCronCommand(serviceName string, spec string, cmd *Command, holdTime time.Duration) {
	root := c.Root()

	if root.CronClient == nil {
		// 初始化
		root.CronClient = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}

	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type:        "distributed-cron",
		Cmd:         cmd,
		Spec:        spec,
		ServiceName: serviceName,
	})

	dirServer := root.GetContainer().MustMake(contract.DirectoryKey).(contract.DirectoryInf)
	distributedServer := root.GetContainer().MustMake(contract.DistributedKey).(contract.Distributed)
	appId := dirServer.AppID()

	var cronCmd Command
	ctx := root.Context()
	cronCmd = *cmd
	cronCmd.args = []string{}
	cronCmd.SetParentNull()

	root.CronClient.AddFunc(spec, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("time error %v", err)
			}
		}()

		selectAppID, err := distributedServer.Select(serviceName, appId, holdTime)
		if err != nil {
			return
		}

		if selectAppID != appId {
			// 证明自己没有被选到
			log.Printf("appid not select %v", appId)
			return
		}

		// 说明当前的appid被选到了
		err = cronCmd.ExecuteContext(ctx)
		if err != nil {
			log.Printf("ExecuteContext error %v", err)
			return
		}
	})
}
