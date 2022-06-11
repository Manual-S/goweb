// Package framework 前缀树的实现
package framework

import (
	"errors"
	"log"
	"strings"
)

type node struct {
	isLast  bool                // 当前的节点是不是最后一个节点
	segment string              // 根据url中的字符串进行分割的segment
	handler []ControllerHandler // 中间件+控制器
	child   []*node             // 当前节点的子节点
	parent  *node               // 指向父节点
}

func NewNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		child:   []*node{},
	}
}

type Tree struct {
	root *node
}

func NewTree() *Tree {
	root := NewNode()
	return &Tree{
		root: root,
	}
}

// isGeneralSegment 是否是一个通用的前缀
func isGeneralSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

//filterChildNodes 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if n.child == nil {
		return nil
	}

	nodes := make([]*node, 0)

	if isGeneralSegment(segment) {
		// 如果是通用路由 则所有的子节点都应该被返回
		return n.child
	}

	for _, v := range n.child {
		if isGeneralSegment(v.segment) {
			// 下一层结点是通用路由 则当前层的所有结点都应该被返回
			nodes = append(nodes, v)
		} else if v.segment == segment {
			nodes = append(nodes, v)
		}
	}

	return nodes
}

// matchNode 返回与uri相匹配的树节点
func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	// 第一部分用来匹配下层结点
	segment := segments[0]

	nodes := n.filterChildNodes(segment)
	if nodes == nil || len(nodes) == 0 {
		// 说明当前的uri一定不在路由树中
		return nil
	}

	// 如果只有一个segment
	if len(segments) == 1 {
		for _, v := range nodes {
			if v.isLast == true {
				return v
			}
		}

		return nil
	}
	// 如果有两个
	for _, cn := range nodes {
		tempNode := cn.matchNode(segments[1])
		if tempNode != nil {
			return tempNode
		}
	}
	return nil
}

func (n *node) parseParamsFromEndNode(uri string) map[string]string {
	ret := make(map[string]string)
	segments := strings.Split(uri, "/")
	cur := n
	cnt := len(segments)

	for i := cnt - 1; i >= 0; i-- {
		if cur.segment == "" {
			break
		}

		if isGeneralSegment(cur.segment) {
			// 是一个通配符
			// cur.segment 是一个:id的结构
			ret[cur.segment[1:]] = segments[i]
		}
		cur = cur.parent
	}

	return ret
}

// AddRouter 向路由树中增加一个路由
func (t *Tree) AddRouter(uri string, handler ...ControllerHandler) error {
	root := t.root

	if root.matchNode(uri) != nil {
		// 说明当前路由树中已经添加过该路由
		return errors.New("router exist")
	}

	// 将uri按照/进行分割
	segments := strings.Split(uri, "/")

	for index, segment := range segments {
		isLast := false

		if index == len(segments)-1 {
			isLast = true
		}

		var objNode *node
		childNodes := root.filterChildNodes(segment)

		if len(childNodes) > 0 {
			// 说明有匹配的子节点
			for _, v := range childNodes {
				if v.segment == segment {
					objNode = v
					break
				}
			}
		}

		if objNode == nil {
			// 没有匹配的子节点 需要创建新的结点
			nNode := NewNode()
			nNode.segment = segment
			if isLast {
				// 是叶子结点
				nNode.isLast = isLast
				nNode.handler = handler
			}

			root.child = append(root.child, nNode)

			nNode.parent = root

			objNode = nNode
		}

		root = objNode
	}

	return nil
}

// FindHandler 根据uri找到对应的处理函数
func (t *Tree) FindHandler(uri string) *node {
	node := t.root.matchNode(uri)
	if node == nil {
		log.Printf("not FindHandler")
		return nil
	}

	return node
}
