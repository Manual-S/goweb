package framework

import (
	"encoding/json"
	"net/http"
)

// Json 输出
func (c *Context) Json(obj interface{}) IResponse {
	if c.HasTimeout() {
		// 已经超时了
		return nil
	}
	c.responseWriter.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(obj)
	if err != nil {
		c.responseWriter.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	c.responseWriter.Write(data)
	return c
}

// Jsonp 输出
func (c *Context) Jsonp(obj interface{}) IResponse {
	return nil
}

// Xml xml 输出
func (c *Context) Xml(obj interface{}) IResponse {
	return nil
}

// Html html 输出
func (c *Context) Html(template string, obj interface{}) IResponse {
	return nil
}

// Text string
func (c *Context) Text(format string, values ...interface{}) IResponse {
	return nil
}

// Redirect 重定向
func (c *Context) Redirect(path string) IResponse {
	return nil
}

// SetHeader header
func (c *Context) SetHeader(key string, val string) IResponse {
	return nil
}

// Cookie
func (c *Context) SetCookie(key string,
	val string,
	maxAge int,
	path,
	domain string,
	secure, httpOnly bool) IResponse {
	return nil
}

// SetStatus 设置状态码
func (c *Context) SetStatus(code int) IResponse {
	return nil
}

// SetOkStatus 设置 200 状态
func (c *Context) SetOkStatus() IResponse {
	c.responseWriter.WriteHeader(http.StatusOK)
	return c
}
