package gone

import (
	"log"
	"net/http"
	"path"
)

/*分组控制(Group Control)*/

type RouteGroup struct {
	prefix      string       // 前缀
	middlewares []HandleFunc // 中间件
	parent      *RouteGroup  // 父分组
	engine      *Engine      // 服务引擎
}

// 一个组创建一个新组
func (group *RouteGroup) Group(prefix string) *RouteGroup {
	engine := group.engine
	// 创建新组
	newGroup := &RouteGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	//
	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

// 添加中间件
func (group *RouteGroup) Use(middlewares ...HandleFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouteGroup) addRoute(method string, comp string, handler HandleFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouteGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandleFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.WriteStatus(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (group *RouteGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}

func (group *RouteGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouteGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
}
