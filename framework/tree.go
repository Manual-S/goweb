// Package framework 前缀树的实现
package framework

import "strings"

type node struct {
	isLast  bool              // 当前的节点是不是最后一个节点
	segment string            // 根据url中的字符串进行分割的segment
	handler ControllerHandler // 自定义的路由处理函数
	child   []*node           // 当前节点的子节点
}

type Tree struct {
	root *node
}

// isGeneralSegment 是否是一个通用的前缀
func isGeneralSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

//filterChildNodes 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	return nil
}

// matchNode 返回与uri相匹配的树节点
func (n *node) matchNode(uri string) *node {
	return nil
}

func NewTree() *Tree {
	root := node{}
	return &Tree{
		root: &root,
	}
}

func (t *Tree) AddRouter(uri string, handler ControllerHandler) error {
	return nil
}

// FindHandler 根据uri找到对应的处理函数
func (t *Tree) FindHandler(uri string) ControllerHandler {
	return nil
}
