package Trees

import (
	"reflect"
	"testing"
)

//
var testValues []int = []int{44, 22, 33, 11, 23, 14, 15, 24}
var sortedValues []int = []int{11,14,15,22,23,24,33,44}
func TestInsert(t *testing.T){

	var BST BinarySearchTree

	
	for index, test := range testValues{
		BST.Insert(test)
	
		if BST.Size != index + 1 {
			t.Errorf("Size is not being updated correctly, Expected: %v, got %v", index +1, BST.Size)
		} 

		_, node, err := BST.Find(test)

		if err != nil{
			t.Error(err)
		}

		if node == nil{
			t.Error("node does not exist in binary search tree")
		}

	}

	if BST.Size != len(testValues){
		t.Error("Size not being incremented correctly")
	}

	for index, test := range sortedValues{
		valIndex, node, err := BST.Find(test)

		if err != nil{
			t.Error(err)
		}

		if sortedValues[valIndex] != sortedValues[index]{
			t.Errorf("Not inserting in sorted order, Expected: %v, got :%v",sortedValues[index],node.Data )
		}
	}

}


func TestFind(t *testing.T){

	var BST BinarySearchTree

	for _, test :=range testValues{
		BST.Insert(test)
	}

	for index, test := range sortedValues{
		valIndex, _ , err := BST.Find(test)

		if err != nil{
			t.Error(err)
		}

		if sortedValues[valIndex] != sortedValues[index]{
			t.Errorf("Binary search is not being indexed correctly, Expected: %v , got %v", sortedValues[index], sortedValues[valIndex])
		}
	}
}

func TestInOrderTraversal(t *testing.T){
	var nodeTestValues []int = []int{44, 22, 33, 11, 23, 14, 15, 24}
	var inOrderValues []int = []int{11,14,15,22,23,24,33,44}

	t.Run("test basic node InOrderTraversal functionality", func(t *testing.T){

		var BSN BinarySearchNode

		for _, test := range testValues{
			BSN.Insert(test)
		}

		xbsnd := BSN.InOrderTraversal(&BSN)
		
		if reflect.DeepEqual(inOrderValues, xbsnd) {
			t.Errorf("basic node traversal not working correctly, Expected %v, got %v", nodeTestValues, xbsnd)
		}
	})

	var BST = &BinarySearchTree{}

	for _, test :=range testValues{
		BST.Insert(test)
	}

	traversalSlice, err := BST.InOrderTraversal(*BST.Root)

	if err != nil{
		t.Error(err)
	}

	if reflect.DeepEqual(traversalSlice, nodeTestValues){
		t.Errorf("BST node traversal is not working correctly, Expected %v, got %v", nodeTestValues, traversalSlice)
	}
}

func TestPreOrderTraversal(t *testing.T){
	var nodeTestValues []int = []int{1,11,14,15,22,23,24,33,44}
	t.Run("test basic node PreOrderTraversal functionality", func(t *testing.T){

		var BSN BinarySearchNode


		for _, test := range testValues{
			BSN.Insert(test)
		}

		xbsn := make([]int, len(testValues))

		BSN.PreOrderTraversal(BSN, xbsn)
		
		if reflect.DeepEqual(nodeTestValues, xbsn) {
			t.Errorf("basic node traversal not working correctly, Expected %v, got %v", nodeTestValues, xbsn)
		}
	})

	var BST = &BinarySearchTree{}

	for _, test :=range testValues{
		BST.Insert(test)
	}

	traversalSlice := BST.PreOrderTraversal(*BST.Root)

	if reflect.DeepEqual(traversalSlice, nodeTestValues){
		t.Errorf("BST node traversal is not working correctly, Expected %v, got %v", nodeTestValues, traversalSlice)
	}
}


func TestPostOrderTraversal(t *testing.T){
	var nodeTestValues []int = []int{15, 14, 11, 24, 23, 33, 22, 44}
	t.Run("test basic node PostOrderTraversal functionality", func(t *testing.T){

		var BSN BinarySearchNode


		for _, test := range testValues{
			BSN.Insert(test)
		}

		xbsn := make([]int, len(testValues))

		BSN.PostOrderTraversal(BSN, xbsn)
		
		if reflect.DeepEqual(nodeTestValues, xbsn) {
			t.Errorf("basic node traversal not working correctly, Expected %v, got %v", nodeTestValues, xbsn)
		}
	})

	var BST = &BinarySearchTree{}

	for _, test :=range testValues{
		BST.Insert(test)
	}

	traversalSlice, err := BST.PostOrderTraversal(*BST.Root)

	if err != nil{
		t.Error(err)
	}

	if reflect.DeepEqual(traversalSlice, nodeTestValues){
		t.Errorf("BST node traversal is not working correctly, Expected %v, got %v", nodeTestValues, traversalSlice)
	}
}

func TestBinarySearch(t *testing.T){

	t.Run("test if it can return if truthy", func(t *testing.T) {
		index := BinarySearch(testValues, 11, 0, len(testValues)-1)

		if index < 0{
			t.Errorf("value should be present in the array, Expected %v got %v",3, index)
		}

		if index != 3{
			t. Errorf("value is not returning the appropriate index, Expected %v got %v", 3, index)
		}
	})
	
	t.Run("test if it can return if not present in the array", func(t *testing.T){
		index := BinarySearch(testValues, 100, 0, len(testValues)-1)

		if index >0{
			t.Errorf("appropraite index is not being returned, Expected %v, got %v", -1, index)
		}
	})
}