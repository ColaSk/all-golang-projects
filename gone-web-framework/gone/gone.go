package gone

import (
	"net/http"
)

// 定义处理函数类型
type HandleFunc func(*Context)

// 定义服务处理引擎
type Engine struct {
	router *router // 定义路由
}

// func (e *Engine) createRouteKey(method string, pattern string) string {
// 	return method + "-" + pattern
// }

func (e *Engine) getRequstRouteKey(req *http.Request) string {
	return req.Method + "-" + req.URL.Path
}

func (e *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// 实现 http.Handler 接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)
	e.router.handle(context)

}

// 验证接口实现
var _ http.Handler = (*Engine)(nil)

// 定义服务处理引擎构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}
