package gee

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	return nil
}
