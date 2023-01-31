package gone

import (
	"net/http"
	"strings"
)

// 定义处理函数类型
type HandleFunc func(*Context)

// 定义服务处理引擎
type Engine struct {
	*RouteGroup         // 继承路由分组，将engine作为顶层分组，并具有分组的一切能力
	router      *router // 定义路由
	groups      []*RouteGroup
}

// func (e *Engine) addRoute(method string, pattern string, handler HandleFunc) {
// 	e.router.addRoute(method, pattern, handler)
// }

// func (e *Engine) GET(pattern string, handler HandleFunc) {
// 	e.addRoute("GET", pattern, handler)
// }

// func (e *Engine) POST(pattern string, handler HandleFunc) {
// 	e.addRoute("POST", pattern, handler)
// }

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// 实现 http.Handler 接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandleFunc

	// 获取请求分组中的中间件
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	context := NewContext(w, req)
	// 将中间件赋值给上下文
	context.handlers = middlewares
	e.router.handle(context)

}

// 验证接口实现
var _ http.Handler = (*Engine)(nil)

// 定义服务处理引擎构造函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
