package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*Node
	handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func newRouter() *Router {
	return &Router{
		roots:    make(map[string]*Node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

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
		parts := parsePattern(node.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
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
		context.handlers = append(context.handlers, router.handlers[key])
	} else {
		context.handlers = append(context.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	context.Next()
}
