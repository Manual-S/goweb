// Package middleware  中间件
package middleware

import (
	"goweb/framework"
	"log"
)

func Test1() framework.ControllerHandler {
	return func(c *framework.Context) error {
		log.Printf("Test1 middlerware start")
		c.Next()
		log.Printf("Test1 middleware end")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	return func(c *framework.Context) error {
		log.Printf("Test2 middlerware start")
		c.Next()
		log.Printf("Test2 middleware end")
		return nil
	}
}
