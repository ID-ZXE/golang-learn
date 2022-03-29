package gee

import "strings"

type Node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*Node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (node *Node) matchChild(part string) *Node {
	for _, child := range node.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (node *Node) matchChildren(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range node.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (node *Node) insert(pattern string, parts []string, height int) {
	// 叶子节点
	if len(parts) == height {
		node.pattern = pattern
		return
	}

	part := parts[height]
	child := node.matchChild(part)
	if child == nil {
		child = &Node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		node.children = append(node.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (node *Node) search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(node.part, "*") {
		if node.pattern == "" {
			return nil
		}
		return node
	}

	part := parts[height]
	children := node.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}