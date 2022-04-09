package frame

import "strings"

type Node struct {
	pattern  string  // url路径
	part     string  // 路由中的一部分，例如 :lang
	children []*Node // 当前节点的子节点
	isWild   bool    // 是否模糊匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点, 用于插入
func (node *Node) matchChild(part string) *Node {
	for _, child := range node.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点, 用于查找
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
	if len(parts) == height {
		node.pattern = pattern
		return
	}

	part := parts[height]
	// 找到第一个匹配成功的节点进行插入
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
