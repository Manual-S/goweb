package demo

import (
	"goweb/framework/cobra"
	"log"
)

func InitFoo() *cobra.Command {
	FooCmd.AddCommand(Foo1Cmd)
	return FooCmd
}

var FooCmd = &cobra.Command{
	Use:   "Foo",
	Short: "定时任务Demo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var Foo1Cmd = &cobra.Command{
	Use:   "foo1",
	Short: "foo1的短说明",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := cmd.GetContainer()
		log.Print("foo1 hello\n")
		log.Printf("Container is %v\n", c)
		return nil
	},
}
