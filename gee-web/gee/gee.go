package gee

import (
	"log"
	"net/http"
)

// 定义请求处理程序
type HandlerFunc func(*Context)

// 分组控制
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // 中间件
	parent      *RouterGroup
	engine      *Engine
}

// 定义服务引擎
type Engine struct {
	*RouterGroup                // 继承
	router       *router        // 路由映射表
	groups       []*RouterGroup // 存储所有组
}

// 创建引擎函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

/*
定义引擎对象功能
*/
// // 添加路由
// func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
// 	engine.router.addRoute(method, pattern, handler)
// }

// // 定义get方法
// func (engine *Engine) GET(pattern string, handler HandlerFunc) {
// 	engine.addRoute("GET", pattern, handler)
// }

// // 定义post方法
// func (engine *Engine) POST(pattern string, handler HandlerFunc) {
// 	engine.addRoute("POST", pattern, handler)
// }

//定义运行服务
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// 实现了Handler接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

/*
定义路由组功能
*/

// 添加组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// 添加路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
