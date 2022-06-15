// Package framework 服务容器
package framework

import (
	"errors"
	"sync"
)

type Container interface {
	// Bind 根据名字绑定一个服务提供者
	Bind(provider ServiceProvider) error
	IsBind(key string) bool
	Make(key string) (interface{}, error)
	MustMake(key string) interface{}
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// WebContainer 具体的服务实现
type WebContainer struct {
	Container
	provider map[string]ServiceProvider
	instance map[string]interface{}
	lock     sync.RWMutex
}

//Bind 根据名字绑定一个服务提供者
func (w *WebContainer) Bind(provider ServiceProvider) error {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.provider[provider.Name()] = provider
	if provider.IsDefer() == false {
		// 说明需要实例化
		method := provider.Register(w)
		params := provider.Params(w)
		ins, err := method(params...)
		if err != nil {
			return err
		}
		// ins是一个由服务提供者初始化的实例
		w.instance[provider.Name()] = ins
	}
	return nil
}

// IsBind 判断key对应的服务提供者是否已经绑定
func (w *WebContainer) IsBind(key string) bool {
	return false
}

func (w *WebContainer) findServiceProvider(key string) ServiceProvider {
	w.lock.RLock()
	defer w.lock.Unlock()
	provider, ok := w.provider[key]
	if ok {
		return provider
	}
	return nil
}

func (w *WebContainer) newInstance(p ServiceProvider, params []interface{}) (interface{}, error) {
	err := p.Boot(w)
	if err != nil {
		return nil, nil
	}

	if params == nil {
		// 获取参数
		params = p.Params(w)
	}

	method := p.Register(w)
	ins, err := method(params...)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (w *WebContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	w.lock.RLock()
	defer w.lock.RUnlock()

	// 查询是否已经注册过服务提供者 如果没有注册 返回错误
	provider := w.findServiceProvider(key)
	if provider == nil {
		return nil, errors.New("provider is nil")
	}
	if forceNew {
		// 强制初始化
		// todo
		return nil, nil
	}

	if ins, ok := w.instance[key]; ok {
		return ins, nil
	}

	// 容器中没有完成实例化
	ins, err := w.newInstance(provider, params)
	if err != nil {
		return nil, err
	}

	w.instance[key] = ins

	return ins, nil
}

// Make 提供获取服务实例的方法
func (w *WebContainer) Make(key string) (interface{}, error) {
	return w.make(key, nil, false)
}
func (w *WebContainer) MustMake(key string) interface{} {
	ins, err := w.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return ins
}
func (w *WebContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return w.make(key, params, true)
}
