package main

import (
	"testing"
)

func TestHeapSort1(t *testing.T) {
	testCase := []int{8, 5, 6, 4, 1, 9, 0, 2, 3, 7}
	nodes := createNodes(testCase)
	var result []int
	result = heapSort(nodes[0], result)
	testFailureFlg := false
	expectedResult := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, e := range expectedResult {
		if result[i] != e {
			testFailureFlg = true
		}
	}
	if testFailureFlg {
		t.Fatalf("test1 failed")
	}
}

func TestHeapSort2(t *testing.T) {
	testCase := []int{8, 5, 5, 6, 4, 1, 9, 0, 2, 3, 4, 7}
	nodes := createNodes(testCase)
	var result []int
	result = heapSort(nodes[0], result)
	testFailureFlg := false
	expectedResult := []int{0, 1, 2, 3, 4, 4, 5, 5, 6, 7, 8, 9}
	for i, e := range expectedResult {
		if result[i] != e {
			testFailureFlg = true
		}
	}
	if testFailureFlg {
		t.Fatalf("test1 failed")
	}
}
