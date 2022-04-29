package gee

import (
	"log"
	"net/http"
	"strings"
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

func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

/*
定义引擎对象功能
*/

//定义运行服务
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// 实现了Handler接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc // 生成一个中间件函数数组

	// 遍历路由组, 获取指定组的中间件
	for _, group := range engine.groups {

		// 判断前缀
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	// 生成上下文
	c := newContext(w, req)
	// 赋值上下文中间件处理函数
	c.handlers = middlewares
	engine.router.handle(c)
}

/*
定义路由组功能
*/

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

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
