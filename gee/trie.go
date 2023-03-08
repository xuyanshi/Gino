package gee

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

/*
插入功能递归查找每一层的节点，如果没有匹配到当前 part 的节点，则新建一个。
需要注意，/p/:lang/doc 只有在第三层节点，即 doc 节点， pattern 才会设置为 /p/:lang/doc。
p 和 :lang 节点的 pattern 属性均为空。
因此，当匹配结束时，使用 n.pattern == "" 来判断路由规则是否匹配成功。
例如，/p/python 虽能成功匹配到 :lang，但 :lang 的 pattern 值为空，因此匹配失败。
*/
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		// Create a new node.
		child = &node{part: part, isWild: part[0] == '*' || part[0] == ':'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

/*
查询功能，同样也是递归查询每一层的节点。
退出规则是，匹配到了*，匹配失败，或者匹配到了第len(parts)层节点。
*/
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(parts[height], "*") {
		if n.pattern == "" {
			// matching failed
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
