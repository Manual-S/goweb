package command

import (
	"errors"
	"fmt"
	"goweb/framework/cobra"
	"goweb/framework/contract"
	"net/http"
)

var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务控制命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个web服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		if container == nil {
			fmt.Printf("container = %v\n", container)
			return errors.New("container is nil")
		}
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		core := kernelService.HttpEngine()
		server := http.Server{
			Handler: core,
			Addr:    ":8080",
		}
		server.ListenAndServe()
		return nil
	},
}

func initAppCommand() *cobra.Command {
	appCommand.AddCommand(appStartCommand)
	return appCommand
}
