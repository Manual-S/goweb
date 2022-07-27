package command

import (
	"fmt"
	"log"

	"goweb/framework/cobra"
	"goweb/framework/contract"
)

func initEnvCmd() *cobra.Command {
	return envCmd
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "获取当前的APP环境",
	Run: func(c *cobra.Command, args []string) {
		container := c.GetContainer()
		if container == nil {
			log.Printf("envCmd con is nil")
			return
		}
		envInf := container.MustMake(contract.EnvKey)
		if envInf == nil {
			fmt.Printf("envInf con is nil")
			return
		}
		envService := envInf.(contract.EnvInf)
		fmt.Println(envService.AppEnv())
	},
}
