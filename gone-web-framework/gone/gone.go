package gone

import (
	"fmt"
	"net/http"
)

// 定义处理函数类型
type HandleFunc func(http.ResponseWriter, *http.Request)

// 定义服务处理引擎
type Engine struct {
	router map[string]HandleFunc // 定义路由
}

func (this *Engine) createRouteKey(method string, pattern string) string {
	return method + "-" + pattern
}

func (this *Engine) getRequstRouteKey(req *http.Request) string {
	return req.Method + "-" + req.URL.Path
}

func (this *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := this.createRouteKey(method, pattern)
	this.router[key] = handler
}

func (this *Engine) GET(pattern string, handler HandleFunc) {
	this.addRoute("GET", pattern, handler)
}

func (this *Engine) POST(pattern string, handler HandleFunc) {
	this.addRoute("POST", pattern, handler)
}

func (this *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, this)
}

// 实现 http.Handler 接口
func (this *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := this.getRequstRouteKey(req)
	if hander, ok := this.router[key]; ok {
		hander(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// 验证接口实现
var _ http.Handler = (*Engine)(nil)

// 定义服务处理引擎构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}
