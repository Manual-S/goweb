package gin

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"strconv"

	"github.com/spf13/cast"
)

// IRequest 代表请求包含的方法
type IRequest interface {
	// 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中带的参数
	// 形如 /book/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	// form表单中带的参数
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

func (ctx *Context) QueryAll() map[string][]string {
	ctx.initQueryCache()
	return ctx.queryCache
}

func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	hash := ctx.QueryAll()
	values, ok := hash[key]
	if ok || len(values) > 0 {
		return cast.ToInt(values[0]), true
	}

	return def, false
}

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	hash := ctx.QueryAll()
	values, ok := hash[key]
	if ok || len(values) > 0 {
		return cast.ToInt64(values[0]), true
	}

	return def, false
}
func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	hash := ctx.QueryAll()
	values, ok := hash[key]
	if ok || len(values) > 0 {
		return cast.ToFloat64(values[0]), true
	}

	return def, false
}
func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	hash := ctx.QueryAll()
	values, ok := hash[key]
	if ok || len(values) > 0 {
		return cast.ToFloat32(values[0]), true
	}

	return def, false
}
func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	hash := ctx.QueryAll()
	values, ok := hash[key]
	if ok || len(values) > 0 {
		return cast.ToBool(values[0]), true
	}

	return def, false
}
func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	hash := ctx.QueryAll()
	values, ok := hash[key]
	if ok || len(values) > 0 {
		return cast.ToString(values[0]), true
	}

	return def, false
}

func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	return nil, false
}

func (ctx *Context) WebParam(key string) interface{} {
	// 调用Gin的Params.Get(name string)函数
	val, ok := ctx.Params.Get(key)
	if ok {
		return val
	}

	return nil
}

func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	val := ctx.WebParam(key)
	if val == nil {
		return def, false
	}

	return cast.ToInt(val), true
}
func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	return 0, false
}
func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	return 0, false
}
func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	return 0, false
}
func (ctx *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	return false, false
}
func (ctx *Context) DefaultParamString(key string, def string) (string, bool) {
	return "", false
}
func (ctx *Context) DefaultParam(key string) interface{} {
	return nil
}

// form 表单中带的参数

func (ctx *Context) DefaultFormInt(key string, def int) (int, bool) {
	mapRes := ctx.FormAll()
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
func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	return 0, false
}
func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	return 0, false
}
func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	return 0, false
}
func (ctx *Context) FormBool(key string, def bool) (bool, bool) {
	return false, false
}
func (ctx *Context) FormString(key string, def string) (string, bool) {
	return "", false
}
func (ctx *Context) FormStringSlice(key string, def []string) ([]string, bool) {
	return nil, false
}

func (ctx *Context) Form(key string) interface{} {
	return nil
}
func (ctx *Context) FormAll() map[string][]string {
	if ctx.Request != nil {
		return map[string][]string(ctx.Request.PostForm)
	}

	return map[string][]string{}
}

// BindJson json body
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.Request != nil {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			return err
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

func (ctx *Context) Uri() string {
	return ctx.Request.RequestURI
}
func (ctx *Context) Method() string {
	return ctx.Request.Method
}
func (ctx *Context) Host() string {
	return ""
}
func (ctx *Context) ClientIp() string {
	return ""
}
func (ctx *Context) Cookies() map[string]string {
	return nil
}
