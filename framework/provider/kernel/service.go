package kernel

import (
	"errors"
	"goweb/framework/contract"
	"goweb/framework/gin"
	"net/http"
)

type KernelService struct {
	contract.Kernel

	engine *gin.Engine
}

func NewKernelService(params ...interface{}) (interface{}, error) {
	if len(params) < 1 {
		return nil, errors.New("params error")
	}
	eng := params[0].(*gin.Engine)
	return &KernelService{engine: eng}, nil
}

func (k *KernelService) HttpEngine() http.Handler {
	return k.engine
}
