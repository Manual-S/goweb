package demo

import (
	demoService "goweb/app/provider/demo"
	"goweb/framework/gin"
	"net/http"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demoService.DemoServiceProvider{})

	r.GET("/demo/demo", api.Demo)

	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{
		service: service,
	}
}

// Demo handler层
func (d *DemoApi) Demo(c *gin.Context) {
	user := d.service.GetUsers()
	c.JSON(http.StatusOK, user)
}
