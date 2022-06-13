package gin

import (
	"encoding/json"
	"log"
	"net/http"
)

// IResponse IResponse代表返回方法
type IResponse interface {
	// IJson Json输出
	IJson(obj interface{}) IResponse

	// IJsonp Jsonp输出
	IJsonp(obj interface{}) IResponse

	// IXml xml输出
	IXml(obj interface{}) IResponse

	// IHtml html输出
	IHtml(template string, obj interface{}) IResponse

	// IText string
	IText(format string, values ...interface{}) IResponse

	// IRedirect 重定向
	IRedirect(path string) IResponse

	// ISetHeader header
	ISetHeader(key string, val string) IResponse

	// ISetCookie Cookie
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// ISetStatus 设置状态码
	ISetStatus(code int) IResponse

	// ISetOkStatus 设置200状态
	ISetOkStatus() IResponse
}

// ISetHeader header
func (ctx *Context) ISetHeader(key string, val string) IResponse {
	ctx.Writer.Header().Add(key, val)
	return ctx
}

func (ctx *Context) IJson(obj interface{}) IResponse {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf("json.Marshal error %v", err)
		return nil
	}
	ctx.ISetHeader("Content-Type", "application/json")
	ctx.Writer.Write(data)
	return ctx
}

func (ctx *Context) ISetOkStatus() IResponse {
	ctx.Writer.WriteHeader(http.StatusOK)
	return ctx
}
