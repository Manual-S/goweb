package command

import "goweb/framework/cobra"

var isCronDaemon = false

func initCronCmd() {
	cronStartCmd.Flags().BoolVarP(&isCronDaemon, "daemon", "d",
		false, "是否以守护进程形式启动一个定时任务")

	cronCmd.AddCommand(cronStartCmd)
}

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "定时任务相关命令", // 短描述 二级命令
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}

		return nil
	},
}

var cronStartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动一个定时任务 常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {

		if isCronDaemon {
			// 以守护进程的形式开启一个定时任务
		}
		cmd.Root().CronClient.Run()
		return nil
	},
}
