package demo

import (
	"fmt"
	"goweb/framework/cobra"
	"log"
)

func InitFoo() *cobra.Command {
	FooCmd.AddCommand(Foo1Cmd)
	return FooCmd
}

var FooCmd = &cobra.Command{
	Use:   "Foo",
	Short: "框架提供的任务Demo",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("this is a command")
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
