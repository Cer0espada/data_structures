package Trees

import (
	"fmt"
	"math"
)

type tree interface {
	search(interface{}) (int, error)
	find(interface{})(int)
	insert(int)
}

type BSN struct{
	left *BSN
	right *BSN
	insert(int)
	find(int)
}

type BinarySearchNode struct {
	Parent *BinarySearchNode
	Left *BinarySearchNode
	Right *BinarySearchNode
	Data int
	Index int
}

func (BSN *BinarySearchNode) hasLeftChild()(bool){
	if BSN.Left != nil{
		return true
	}
	return false
}

func (BSN *BinarySearchNode) hasRightChild()(bool){
	if BSN.Right != nil{
		return true
	}
	return false
}

func (BSN *BinarySearchNode) Insert(value int){
	
	if BSN.Data > value {
		if !BSN.hasLeftChild(){
			BSN.Left = &BinarySearchNode{Data:value, Parent:BSN}
			return
		}
		BSN.Left.Insert(value)
		return
	}

	if !BSN.hasRightChild(){
		BSN.Right = &BinarySearchNode{Data:value, Parent: BSN}
		return
	}
	BSN.Right.Insert(value)
}

func (BSN *BinarySearchNode) Find(value int, index *int,node *BinarySearchNode){
	if BSN.Data == 0 || BSN == nil{
		return 
	}

	if BSN.Data == value {
		*index = BSN.Index
		*node = *BSN
		return 
	}

	if BSN.Data > value{
		if BSN.Left != nil {
			BSN.Left.Find(value, index,node)
		}
		return
	}

	if BSN.Right != nil{
		BSN.Right.Find(value, index,node)
		return 
	}
}

func (BSN *BinarySearchNode) Delete(node *BinarySearchNode, value int)(*BinarySearchNode){
	if node == nil {
		return node 
	}

	if value > BSN.Data {
		BSN.Right = BSN.Delete(BSN.Right, value)
	}
	if value < BSN.Data {
		BSN.Left = BSN.Delete(BSN.Left, value)
	}

	if BSN.Left == nil && BSN.Right == nil {
		BSN = nil
	}

	if BSN.Right == nil {
		BSN.Data = BSN.successor()
		BSN.Right = BSN.Delete(BSN.Right, BSN.Data)

	}

	BSN.Data = BSN.predecessor()
	BSN.Left = BSN.Delete(BSN.Left, BSN.Data)
	return BSN
}

func (BSN *BinarySearchNode) successor()(int){
	node := *BSN
	node = *node.Right

	for node.Left != nil{
		node = *node.Left
	}
	return node.Data
}

func (BSN *BinarySearchNode) predecessor()(int){
	node := *BSN
	node = *node.Left
	for node.Right != nil{
		node = *node.Right
	}
	return node.Data
}

func (BSN *BinarySearchNode) setIndex(count *int){
	if BSN == nil{
		return
	}
	
	if BSN.hasLeftChild(){
		BSN.Left.setIndex(count)
	}

	index := *count
	*count = *count +1
	BSN.Index = index

	if BSN.hasRightChild(){
		BSN.Right.setIndex(count)
	}

}
func (BSN *BinarySearchNode) InOrderTraversal(node *BinarySearchNode)([]int){
	if node == nil{
		return []int{}
	}
	l := node.InOrderTraversal(node.Left)
	r := node.InOrderTraversal(node.Right)
	l = append(l , node.Data)
	l = append(l, r...)
	return l
}

func (BSN *BinarySearchNode) PreOrderTraversal(node BinarySearchNode, xbsn []int)(){
	
	xbsn[node.Index] = node.Data

	if node.Left != nil{
		node.PreOrderTraversal(*node.Left, xbsn)
	}

	if node.Right != nil {
		node.PreOrderTraversal(*node.Right, xbsn)
	}
}

func (BSN *BinarySearchNode) PostOrderTraversal(node BinarySearchNode, xbsn []int)(){

	if node.Left != nil{
		node.PreOrderTraversal(*node.Left, xbsn)
	}

	if node.Right != nil {
		node.PreOrderTraversal(*node.Right, xbsn)
	}

	xbsn[node.Index] = node.Data
}

// A Binary Search Tree (BST) is a tree in which all the nodes follow the below-mentioned properties −
// The value of the key of the left sub-tree is less than the value of its parent (root) node's key.
// The value of the key of the right sub-tree is greater than or equal to the value of its parent (root) node's key.
// This data structure allows for easy lookup of values when search
type BinarySearchTree struct {
	Root *BinarySearchNode
	Size int 
}

func (BST *BinarySearchTree) isEmpty()(bool){
	if BST.Size == 0 || BST.Root == nil{
		return true
	}
	return false
}

func (BST *BinarySearchTree) hasOne()(bool){
	if BST.Size == 1 && BST.Root != nil{
		return true
	}
	return false
}

func (BST *BinarySearchTree) Insert(value int){
	
	if BST.isEmpty(){
		node := &BinarySearchNode{Data:value}
		BST.Root = node
		node.Index = 0 
		BST.Size++
		return
	}
	
	BST.Root.Insert(value)
	var count int = 0
	BST.Root.setIndex(&count)
	BST.Size++
}
//returns the binarySearchNode at a particular index 
func (BST *BinarySearchTree) Find(value int)(int, *BinarySearchNode, error){
	if BST.isEmpty(){
		err:= fmt.Errorf("this Binary Search Tree is empty")
		return -1, nil,  err
	}

	var node BinarySearchNode
	var index int  = -1
	BST.Root.Find(value,&index, &node)

	if index < 0 {
		err := fmt.Errorf("this item does not exist in the Binary Search Tree")
		return -1, &node, err
	}
	return index, &node, nil
}
//Performs InOrderTraversal of binary Tree
func (BST *BinarySearchTree) InOrderTraversal(node BinarySearchNode)([]int, error){

	if BST.isEmpty(){
		err:= fmt.Errorf("this binary search tree is empty")
		return []int{}, err
	}
	if BST.hasOne() {
		node = *BST.Root
		return []int{BST.Root.Data}, nil
	}

	xbsn := BST.Root.InOrderTraversal(BST.Root)
	return xbsn, nil
}
//Performs PreOrderTraversal of binary Tree
func (BST *BinarySearchTree) PreOrderTraversal(node BinarySearchNode)([]int){
	xbsn := make([]int, BST.Size)

	xbsn[node.Index] = node.Data

	if BST.Root.Left != nil{
		node.PreOrderTraversal(*BST.Root.Left,xbsn)
	}

	if BST.Root.Right != nil {
		node.PreOrderTraversal(*BST.Root.Right, xbsn)
	}
	return xbsn
}
//Performs PostOrderTraversal of binary Tree
func (BST *BinarySearchTree) PostOrderTraversal(node BinarySearchNode)([]int, error){
	if BST.isEmpty(){
		err:= fmt.Errorf("this binary search tree is empty")
		return []int{}, err
	}
	if BST.hasOne() {
		node = *BST.Root
		return []int{BST.Root.Data}, nil
	}

	xbsn := make([]int, BST.Size)

	if BST.Root.Left != nil{
		node.PostOrderTraversal(*BST.Root.Left, xbsn) 
	}

	if BST.Root.Right != nil {
		node.PostOrderTraversal(*BST.Root.Right, xbsn)
	}

	xbsn[node.Index] = node.Data

	return xbsn, nil
}

func (BST *BinarySearchTree) Delete(value int)(error){
	if BST.isEmpty(){
		err:= fmt.Errorf("the tree is empty")
		return err
	}

	if BST.hasOne(){
		BST.Root = nil
		BST.Size = 0 
		return nil
	}

	BST.Delete(value)
	var count int =0
	BST.Root.setIndex(&count)
	return nil
}

func (BSN *BinarySearchNode) getHeight()(int){
	if BSN == nil{
		return -1
	}
	h1 := BSN.Left.getHeight()
	h2 := BSN.Right.getHeight()
	return int(math.Max(float64(h1),float64(h2)) +1 )
}

func (BST *BinarySearchTree) GetHeight()(int){
	if BST.isEmpty() {
		return 0
	}

	if BST.hasOne(){
		return 1
	}

	h := BST.Root.getHeight()
	return h
}
//performs binary search on a slice of ints, conducted in O(log(n))
func BinarySearch (arr []int, value int, start int, end int)(int){
	if start > end {
		return -1
	}
	
	mid := start + end /2

	if arr[mid] == value {
		return mid
	}
	
	if arr[mid] > value{
		BinarySearch(arr,value,start, mid-1)
	}
	
	BinarySearch(arr,value, mid+1, end)

	return -1
}