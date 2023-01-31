package gone

import (
	"gone/trie"
	"log"
	"net/http"
	"strings"
)

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
type router struct {
	handlers map[string]HandleFunc
	roots    map[string]*trie.Node
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
		roots:    make(map[string]*trie.Node),
	}
}

// 解析路由
func parsePattern(pattern string) []string {

	split := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, s := range split {
		if s != "" {
			parts = append(parts, s)
			if strings.HasPrefix(s, "*") {
				break
			}
		}
	}
	return parts

}

func (r *router) createRouteKey(method string, pattern string) string {
	return method + "-" + pattern
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := r.createRouteKey(method, pattern)
	parts := parsePattern(pattern)

	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &trie.Node{}
	}

	root := r.roots[method]
	root.Insert(pattern, parts, 0)

	r.handlers[key] = handler
}

// 获取路由节点与请求的路由参数
func (r *router) getRoute(method string, path string) (*trie.Node, map[string]string) {
	log.Printf("Requst route %4s - %s", method, path)

	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	// 搜索路由节点
	node := root.Search(searchParts, 0)

	// 构建路由中的参数列表
	if node != nil {
		parts := parsePattern(node.GetPattern())

		for index, part := range parts {
			if strings.HasPrefix(part, ":") {
				params[part[1:]] = searchParts[index]
			}

			if strings.HasPrefix(part, "*") {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}

		return node, params

	}

	return nil, nil

}

func (r *router) handle(context *Context) {
	// 获取路由节点与路由中的参数列表
	node, params := r.getRoute(context.Method, context.Path)

	if node != nil {
		context.Params = params // 添加路由参数
		key := r.createRouteKey(context.Method, node.GetPattern())
		context.handlers = append(context.handlers, r.handlers[key])
	} else {
		context.handlers = append(context.handlers, func(context *Context) {
			context.WriteString(
				http.StatusNotFound,
				"404 NOT FOUND: %s\n",
				context.Path)
		})
	}

	context.Next()
}
