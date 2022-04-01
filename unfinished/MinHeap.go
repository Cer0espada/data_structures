package Trees

import (
	"fmt"
	"math"
)

type TreeNode interface {
	insert() int
	hasLeftChild() bool
	hasRightChild() bool
	hasParent() bool
	swap(*TreeNode, *TreeNode)
}

type Tree interface {
	isEmpty() bool
	hasOne() bool
}

type MinHeapNode struct {
	Parent *MinHeapNode
	Left   *MinHeapNode
	Right  *MinHeapNode
	Value  int
	Index  int
}

func (MHN *MinHeapNode) hasLeftChild() bool {
	if MHN.Left != nil {
		return true
	}
	return false
}

func (MHN *MinHeapNode) hasRightChild() bool {
	if MHN.Right != nil {
		return true
	}
	return false
}

func (MHN *MinHeapNode) hasParent() bool {
	if MHN.Parent != nil {
		return true
	}
	return false
}

func (MHN *MinHeapNode) Insert(value int) {

	if !MHN.hasLeftChild() {
		MHN.Left = &MinHeapNode{Parent: MHN, Value: value}
		return
	}

	if !MHN.hasRightChild() {
		MHN.Right = &MinHeapNode{Parent: MHN, Value: value}
		return
	}

	if MHN.hasLeftChild(){
		MHN.Left.Insert(value)
		return
	}

	if MHN.hasRightChild(){
		MHN.Right.Insert(value)
		return
	}
}

func (MHN *MinHeapNode) setIndex(count *int){

	index := *count
	*count = *count +1
	MHN.Index = index

	if MHN.hasLeftChild(){
		MHN.Left.setIndex(count)
	}
	
	if MHN.hasRightChild(){
		MHN.Right.setIndex(count) 
	}
	
}

func (MHN *MinHeapNode) getIndex(index int, node *MinHeapNode){
	if MHN == nil{
        return
    }

    if MHN.Index == index{
        *node = *MHN
        return
    }
        MHN.Left.getIndex(index, node)
        MHN.Right.getIndex(index,node)
}

func (MHN *MinHeapNode) Delete(Value int){

}

type MinHeapTree struct {
	Root *MinHeapNode
	Size int
}

func (MHT *MinHeapTree) isEmpty() bool {
	if MHT.Size == 0 {
		return true
	}
	return false
}

func (MHT *MinHeapTree) hasOne() bool {
	if  MHT.Root.Left == nil && MHT.Root.Right == nil && MHT.Size == 1 {
		return true
	}
	return false
}


func (MHT *MinHeapTree) Peek()(int, error){
	if MHT.isEmpty() {
		err := fmt.Errorf("Tree is empty")
		return -1, err
	}

	return MHT.Root.Value, nil
}

func (MHT *MinHeapTree) Poll()(int, error){
	if MHT.isEmpty(){
		err := fmt.Errorf("Tree is empty")
		return -1, err
	}

	val := MHT.Root.Value
	MHT.HeapifyDown()
	return val, nil
}
func (MHT *MinHeapTree) Insert(value int){

	if MHT.isEmpty(){
		MHT.Root = &MinHeapNode{Value:value, Index: 0}
		MHT.Size++
		return
	}
	//if tree is at maximum capacity for its height make it a left child of the left most node
	if int(MHT.height()*100) % 100 ==0{
		node := MHT.Root

		for node.Left !=nil{
			node = node.Left
		}
		node.Left = &MinHeapNode{Parent: node, Value:value, Index: MHT.Size}
		MHT.HeapifyUp()
		MHT.Size++
		return
	}
	MHT.Root.Insert(value)
	MHT.Size++
	var count int = 0
	MHT.Root.setIndex(&count)

	MHT.HeapifyUp()
}

func (MHT *MinHeapTree) HeapifyUp(){
	index := MHT.Size -1
	
	var node MinHeapNode
	MHT.Root.getIndex(index,&node)

	for node.Parent!= nil && node.Parent.Value > node.Value {

		if node.Parent.Value > node.Value{
			swap(&node.Value, node.Parent.Value)
		}
		node = *node.Parent
	}
}

func swap(first, second interface{}){
	var temp interface{}
	temp = first
	first = second 
	second = temp
}
// calculates the max number of nodes given a height of the tree
//					h = 2^(h+1)-1
func maxNodesForHeight(h int)(int){ 
	return int(math.Pow(float64(h+1),2)) 
}
//calculates the height given the number of nodes in a tree
func (MHT *MinHeapTree) height()(float64){ 
	return math.Log2(float64(MHT.Size + 1)) +1
}

func (MHT *MinHeapTree) HeapifyDown(){

	if MHT.hasOne(){
		MHT = &MinHeapTree{}
		MHT.Size--
		return
	}
	node := MHT.Root
	lastNode, _ := MHT.getIndex(MHT.Size-1)
	node.Value = lastNode.Value
	lastNode = nil
	MHT.Size-- 
	for node.hasLeftChild(){

		//grabbing and assuming the left child is a smaller node
		smallerChildNode := node.Left

		if node.hasRightChild(){
			//if the right happens to be smaller than the left then grab that node 
			if node.Right.Value < node.Left.Value {
				smallerChildNode = node.Right
			}
		}

		//if the current node is smaller than the child node we're done and can break out 
		if node.Value < smallerChildNode.Value{
				break
		}

		//swap that node with the child node and repeat
		if node.Value < smallerChildNode.Value{
			swap(&node, node.Left)
		}
		node = node.Left
	}

}

func (MHT *MinHeapTree) getIndex(index int)(*MinHeapNode, error){
	if MHT.Size < index +1 {
		err := fmt.Errorf("index exceeds tree size")
		return nil, err
	} 
	var node MinHeapNode
	MHT.Root.getIndex(index, &node)
	return &node, nil

}
