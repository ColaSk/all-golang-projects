package gone

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 定义上下文
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string // 路由中解析的参数列表
}

func (c *Context) Param(key string) string {
	value := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) WriteStatus(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) WriteString(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.WriteStatus(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) WriteJSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.WriteStatus(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) WriteData(code int, data []byte) {
	c.WriteStatus(code)
	c.Writer.Write(data)
}

func (c *Context) WriteHTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.WriteStatus(code)
	c.Writer.Write([]byte(html))
}

// 上下文构造函数
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}
