package gee

import (
	"fmt"
	"net/http"
)

// 定义请求处理程序
type HandlerFunc func(http.ResponseWriter, *http.Request)

// 定义服务引擎
type Engine struct {
	router map[string]HandlerFunc // 路由映射表
}

// 创建引擎函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

/*
定义引擎对象功能
*/
// 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	var key string = method + "-" + pattern
	engine.router[key] = handler
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

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var key string = req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
