package main

import "fmt"

type node struct {
	value     int
	leftNode  *node
	rightNode *node
}

func main() {
	testVal := []int{5, 4, 7, 3, 0, 9, 2, 1, 8, 6}
	nodes := createNodes(testVal)
	fmt.Println(nodes)
}

func createNodes(slice []int) []*node {
	nodes := make([]*node, 0)
	for i, s := range slice {
		nodeCreated := node{s, nil, nil}
		var parentNode *node
		if i == 0 {
			parentNode = nil
		} else {
			parentNode = nodes[0]
		}
		appendNode(&nodeCreated, parentNode)
		nodes = append(nodes, &nodeCreated)
	}
	return nodes
}

func appendNode(nodeCreated *node, parentNode *node) {
	if parentNode == nil {
		return
	}
	if nodeCreated.value < parentNode.value {
		if parentNode.leftNode == nil {
			parentNode.leftNode = nodeCreated
		} else {
			appendNode(nodeCreated, parentNode.leftNode)
		}
	} else if nodeCreated.value > parentNode.value {
		if parentNode.rightNode == nil {
			parentNode.rightNode = nodeCreated
		} else {
			appendNode(nodeCreated, parentNode.rightNode)
		}
	} else {
		if parentNode.leftNode == nil {
			parentNode.leftNode = nodeCreated
		} else if parentNode.rightNode == nil {
			parentNode.rightNode = nodeCreated
		} else {
			appendNode(nodeCreated, parentNode.rightNode)
		}
	}

}
