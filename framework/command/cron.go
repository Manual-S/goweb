package command

import (
	"fmt"
	"goweb/framework/cobra"
	"goweb/framework/contract"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var isCronDaemon = false

func initCronCmd() *cobra.Command {
	cronStartCmd.Flags().BoolVarP(&isCronDaemon, "daemon", "d",
		false, "是否以守护进程形式启动一个定时任务")

	cronCmd.AddCommand(cronStartCmd)

	return cronCmd
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
		container := cmd.GetContainer()

		dirService := container.MustMake(contract.DirectoryKey).(contract.DirectoryInf)

		pidFolder := dirService.RuntimeFolder() // pid文件的路径
		serverPidFile := filepath.Join(pidFolder, "cron.pid")

		//logFolder := dirService.LogFolder() // 日志文件的路径
		//logFile := filepath.Join(logFolder, "cron.log")

		//curFolder := dirService.BaseFolder()

		if isCronDaemon {
			// 以守护进程的形式开启一个定时任务
		}

		fmt.Println("start cron job")
		pid := strconv.Itoa(os.Getpid())
		fmt.Println("pid is ", pid)
		err := ioutil.WriteFile(serverPidFile, []byte(pid), 0664)
		if err != nil {
			fmt.Errorf("WriteFile error %v", err)
			return err
		}
		fmt.Println("start corn task")
		cmd.Root().CronClient.Run()
		return nil
	},
}
