package trie

import (
	"strings"
)

/*
Trie 树 - 动态路由(dynamic route)解析
支持两种模式:name和*filepath
*/

type Node struct {
	pattern  string  // 待匹配的路由
	part     string  // 当前节点路由
	children []*Node // 子节点
	isWild   bool    // 是否精准匹配 part 含有 : 或 * 时为true
}

func (n *Node) GetPattern() string {
	return n.pattern
}

// 获取第一个匹配成功的子节点
func (n *Node) matchChildOne(part string) *Node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 获取所有匹配成功的子节点
func (n *Node) matchChildren(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 递归注册节点
func (n *Node) Insert(pattern string, parts []string, height int) {

	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChildOne(part)

	if child == nil {
		isWild := strings.HasPrefix(part, "*") || strings.HasPrefix(part, ":")
		child = &Node{part: part, isWild: isWild}
		n.children = append(n.children, child)
	}

	child.Insert(pattern, parts, height+1)

}

// 递归搜索节点
func (n *Node) Search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	childs := n.matchChildren(part)

	for _, child := range childs {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
