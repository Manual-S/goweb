package framework

import (
	"mime/multipart"
	"net"
	"strconv"
	"strings"

	"github.com/spf13/cast"
)

func (c *Context) QueryInt(key string, def int) (int, bool) {
	mapRes := c.QueryAll()
	values, ok := mapRes[key]
	if ok {
		if len(values) > 0 {
			return cast.ToInt(values[len(values)-1]), true
		}
	}
	return def, false
}
func (c *Context) QueryInt64(key string, def int64) (int64, bool) {
	mapRes := c.QueryAll()
	values, ok := mapRes[key]
	if ok {
		if len(values) > 0 {
			return cast.ToInt64(values[len(values)-1]), true
		}
	}
	return def, false
}
func (c *Context) QueryFloat64(key string, def float64) (float64, bool) {
	mapRes := c.QueryAll()
	values, ok := mapRes[key]
	if ok {
		if len(values) > 0 {
			return cast.ToFloat64(values[len(values)-1]), true
		}
	}
	return def, false
}
func (c *Context) QueryFloat32(key string, def float32) (float32, bool) {
	mapRes := c.QueryAll()
	values, ok := mapRes[key]
	if ok {
		if len(values) > 0 {
			return cast.ToFloat32(values[len(values)-1]), true
		}
	}
	return def, false
}
func (c *Context) QueryBool(key string, def bool) (bool, bool) {
	mapRes := c.QueryAll()
	values, ok := mapRes[key]
	if ok {
		if len(values) > 0 {
			return cast.ToBool(values[len(values)-1]), true
		}
	}
	return def, false
}
func (c *Context) QueryString(key string, def string) (string, bool) {
	mapRes := c.QueryAll()
	values, ok := mapRes[key]
	if ok {
		if len(values) > 0 {
			return cast.ToString(values[len(values)-1]), true
		}
	}
	return def, false
}
func (c *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	return nil, false
}
func (c *Context) Query(key string) interface{} {
	return nil
}
func (c *Context) QueryArray(key string, def []string) []string {
	return nil
}
func (c *Context) QueryAll() map[string][]string {
	if c.request != nil {
		return map[string][]string(c.request.URL.Query())
	}

	return map[string][]string{}
}

// 路由匹配中带的参数
// 形如 /book/:id

func (c *Context) ParamInt(key string, def int) (int, bool) {
	return 0, false
}
func (c *Context) ParamInt64(key string, def int64) (int64, bool) {
	return 0, false
}
func (c *Context) ParamFloat64(key string, def float64) (float64, bool) {
	return 0, false
}
func (c *Context) ParamFloat32(key string, def float32) (float32, bool) {
	return 0, false
}
func (c *Context) ParamBool(key string, def bool) (bool, bool) {
	return false, false
}
func (c *Context) ParamString(key string, def string) (string, bool) {
	return "", false
}
func (c *Context) Param(key string) interface{} {
	return nil
}

// form 表单中带的参数

func (c *Context) FormInt(key string, def int) (int, bool) {
	mapRes := c.FormAll()
	if valus, ok := mapRes[key]; ok {
		if len(valus) > 0 {
			res, err := strconv.Atoi(valus[len(valus)-1])
			if err != nil {
				return def, false
			}

			return res, true
		}
	}

	return def, false
}
func (c *Context) FormInt64(key string, def int64) (int64, bool) {
	return 0, false
}
func (c *Context) FormFloat64(key string, def float64) (float64, bool) {
	return 0, false
}
func (c *Context) FormFloat32(key string, def float32) (float32, bool) {
	return 0, false
}
func (c *Context) FormBool(key string, def bool) (bool, bool) {
	return false, false
}
func (c *Context) FormString(key string, def string) (string, bool) {
	return "", false
}
func (c *Context) FormStringSlice(key string, def []string) ([]string, bool) {
	return nil, false
}
func (c *Context) FormFile(key string) (*multipart.FileHeader, error) {
	return nil, nil
}
func (c *Context) Form(key string) interface{} {
	return nil
}
func (c *Context) FormAll() map[string][]string {
	if c.request != nil {
		return map[string][]string(c.request.PostForm)
	}

	return map[string][]string{}
}

// BindJson json body
func (c *Context) BindJson(obj interface{}) error {
	return nil
}

// BindXml xml body
func (c *Context) BindXml(obj interface{}) error {
	return nil
}

// GetRawData 其他格式
func (c *Context) GetRawData() ([]byte, error) {
	return nil, nil
}

// 基础信息

func (c *Context) Uri() string {
	return c.request.RequestURI
}
func (c *Context) Method() string {
	return c.request.Method
}
func (c *Context) Host() string {
	return c.request.Host
}

// ClientIp 获取ip地址
// todo 获取ip这里有需要http知识
// 参考资料 https://www.cnblogs.com/GaiHeiluKamei/p/13731791.html
func (c *Context) ClientIp() string {
	ip := c.request.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip
	}
	ip = c.request.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i
		}
	}
	if ip == "" {
		ip = c.request.RemoteAddr
	}
	return ip
}

// Headers header
func (c *Context) Headers() map[string][]string {
	return nil
}
func (c *Context) Header(key string) (string, bool) {
	return "", false
}

// Cookies cookie
func (c *Context) Cookies() map[string]string {
	return nil
}
func (c *Context) Cookie(key string) (string, bool) {
	return "", false
}
