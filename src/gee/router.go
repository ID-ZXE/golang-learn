package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*Node
	handlers map[string]HandlerFunc
}

// 小写的方法名代表访问控制在包内
func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*Node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			// 解析到首字母为*之后 停止解析
			// 后续的路径都不再需要 直接模糊匹配
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// GET[/index/:name] name为路径参数
// GET[/index/*] *表示模糊匹配
func (router *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &Node{}
	}
	router.roots[method].insert(pattern, parts, 0)
	router.handlers[key] = handler
}

func (router *Router) getRoute(method string, path string) (*Node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := router.roots[method]

	if !ok {
		return nil, nil
	}

	node := root.search(searchParts, 0)

	if node != nil {
		// 解析构建的node节点的pattern
		parts := parsePattern(node.pattern)
		for index, part := range parts {
			// 如果是:param模式 则将将参数put到params中
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			// 模糊匹配
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return node, params
	}

	return nil, nil
}

func (router *Router) handle(context *Context) {
	node, params := router.getRoute(context.Method, context.Path)

	if node != nil {
		key := context.Method + "-" + node.pattern
		context.Params = params
		// 将route中获取到handler放入处理链条的最末端
		context.handlers = append(context.handlers, router.handlers[key])
	} else {
		context.handlers = append(context.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	context.Next()
}
