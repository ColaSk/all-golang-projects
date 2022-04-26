package gee

import (
	"net/http"
)

// 定义请求处理程序
type HandlerFunc func(*Context)

// 定义服务引擎
type Engine struct {
	router *router // 路由映射表
}

// 创建引擎函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

/*
定义引擎对象功能
*/
// 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// 定义get方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 定义post方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//定义运行服务
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// 实现了Handler接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
