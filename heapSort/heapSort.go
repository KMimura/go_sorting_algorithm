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
	var sortedVal []int
	sortedVal = heapSort(nodes[0], sortedVal)
	fmt.Println(sortedVal)
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

func heapSort(nodeVal *node, sortedSlice []int) []int {
	fmt.Println(nodeVal.value)
	if nodeVal.leftNode != nil {
		sortedSlice = heapSort(nodeVal.leftNode, sortedSlice)
		sortedSlice = append(sortedSlice, nodeVal.value)
	} else {
		sortedSlice = append(sortedSlice, nodeVal.value)
	}
	if nodeVal.rightNode != nil {
		sortedSlice = heapSort(nodeVal.rightNode, sortedSlice)
	}
	return sortedSlice
}
