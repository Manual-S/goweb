package console

import (
	"errors"
	"fmt"
	"goweb/app/console/command/demo"
	"goweb/framework"
	"goweb/framework/cobra"
	"goweb/framework/command"
)

// RunCommand 初始化根command并且执行
func RunCommand(c framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "goweb",
		Short: "goweb提供的命令行工具",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	if c == nil {
		fmt.Printf("RunCommand container is nil\n")
		return errors.New("container is nil")
	}

	rootCmd.SetContainer(c)

	// 为rootCmd增加框架提供的命令
	command.AddKernelCommands(rootCmd)

	// 添加自定义的命令
	AddNormalCmd(rootCmd)

	return rootCmd.Execute()
}

// AddNormalCmd 绑定业务自定义的命令
func AddNormalCmd(rootCmd *cobra.Command) {
	rootCmd.AddCronCommand("* * * * * *", demo.FooCmd)
	//rootCmd.AddCommand(demo.InitFoo())
	// 分布式定时任务
	//rootCmd.AddDistributedCronCommand("foo_func_for_test", "*/5 * * * * *", demo.FooCmd, 2*time.Second)
}
