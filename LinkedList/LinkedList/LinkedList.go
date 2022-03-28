package LinkedList

import (
	"errors"
	"fmt"
)

type ListNode struct{
	Next *ListNode
	Value int
}

//Singly linked lists are the simplest type of linked list
// There is a single track that we can traverse the list in; we start at the head node, 
// traverse from the root until the last node, which will end at an empty null value.
//Linked List tend to be more memory efficent and sacrifice the typical O(1) index look up that Arrays/ Lists /Slices have 
//Insert at N position O(N)
//Insert at end O(N)
//Remove at Front O(1)
//Remove at N position O(N)
//Remove at End O(N)
type LinkedList struct{
	Head *ListNode
	Size int
}

func (LL *LinkedList) isEmpty()(bool){
	if LL.Head == nil && LL.Size ==0{
		return true
	}
	return false
}

func (LL *LinkedList) hasOne()(bool){
	if LL.Head != nil && LL.Size ==1{
		return true
	}
	return false
}
//Insert at Front O(1)
func (LL *LinkedList) AddFront(value int) (error){
	
	if LL.isEmpty() {
		
		LL.Head = &ListNode{Value: value}
		LL.Size++
		return nil
	}

	curr := ListNode{LL.Head, value}
	curr.Next = LL.Head
	LL.Head = &curr
	LL.Size++
	return nil
}
//Adding at middle is O(N) time and is unstable
func(LL *LinkedList) AddMiddle(value int)(error){

	if LL.isEmpty(){
		newNode := ListNode{Value: value}
		LL.Head = &newNode
		LL.Size++
		return nil
	}

	if LL.hasOne(){
		newNode := ListNode{Next: LL.Head, Value: value}
		LL.Head.Next = &newNode
		LL.Size++
		return nil
	}

	curr := LL.Head

	mid := LL.Size / 2 

	for i :=0; i <= mid ;i++{
		if curr.Next == nil{
			err := errors.New("LinkedList pointer error")
			return err
		}

		curr = curr.Next
	}
	
	nextNode := curr.Next

	curr.Next = &ListNode{nextNode, value}

	LL.Size++
	return nil

}

// Adding to the end of the linkedlist is O(N) time without tail optimization
func (LL *LinkedList) AddEnd(value int){
	
	if LL.isEmpty(){
		LL.Head = &ListNode{Value: value}
		LL.Size++
		return 
	}
	curr := LL.Head

	for curr.Next != nil{
		curr = curr.Next
	}

	curr.Next = &ListNode{Value:value}
	LL.Size++
}
//Removing from the beginning of the lsit is O(1) time 
func (LL *LinkedList) RemoveFront()(error) {

	if LL.isEmpty(){
		err := errors.New("Linked List is already empty")
		return err
	}

	if LL.hasOne(){
		LL.Head = nil
		LL.Size--
		return nil
	}

	LL.Head = LL.Head.Next
	LL.Size--
	return nil
}
//Unstable rmeoval of middle element completed in O(N) time
func (LL *LinkedList) RemoveMiddle() (error) {
	if LL.isEmpty(){
		err := errors.New("Linked List is already empty")
		return err
	}

	if LL.hasOne(){
		LL.Head = nil
		LL.Size--
		return nil
	}

	mid := LL.Size /2

	curr := LL.Head

	var prev ListNode

	for i :=0; i< mid;i++{
		prev = *curr
		curr = curr.Next
	}

	nextNode := curr.Next
	prev.Next = nextNode
	LL.Size--
	return nil

}
//Removal of end without tail optimization is completed in O(N) time
func (LL *LinkedList) RemoveEnd()(error){
	if LL.isEmpty(){
		err := errors.New("Linked List is already empty")
		return err
	}

	if LL.hasOne(){
		LL.Head = nil
		LL.Size--
		return nil
	}

	curr := LL.Head

	for{
		if curr.Next == nil{
			break
		}
		curr = curr.Next
	}

	curr = nil
	LL.Size--
	return nil
}

//returns the index of the element in the chain of a linked list O(N)
func (LL *LinkedList) indexOf(value int) (int ,error){
	if LL.isEmpty(){
		err := errors.New("Linked List empty")
		return -1, err
	}

	if LL.hasOne() || LL.Head.Value == value {
		return 0,nil
	}

	curr := LL.Head
	var count int

	for curr.Next != nil{
		if curr.Value == value{
			return count, nil
		}
		count++
		curr = curr.Next
	}

	if curr.Value != value{
		err := fmt.Errorf("value of %v does not exist in the LinkedList", value)
		return -1, err
	}
	
	return count, nil
}

func (LL *LinkedList) DetectCycle()(bool, error){
	if LL.isEmpty() || LL.hasOne() || LL.Size < 2{
		return false, nil
	}

	fast := LL.Head
	slow := LL.Head

	for {
		if slow.Next == nil || fast.Next.Next ==nil {
			return false, nil
		}

		if slow == fast {
			return true, nil
		}

		fast = fast.Next.Next
		slow = slow.Next
	}
	
}

func (LL *LinkedList) ReorderList(step int)(error){

	if LL.isEmpty() || LL.hasOne() || LL.Size < 2{
		err := errors.New(fmt.Sprintf("step size of %v is too small to do any reordering", step))
		return err
	}

	capacity := LL.Size/step

	nodeArray := make([]ListNode,capacity, capacity)

	curr := LL.Head

	var count int
	var prev ListNode

	for {
		if curr.Next == nil{
			break
		}

		if count % step ==0 {
			nodeArray[count] = *curr
			prev.Next = curr.Next
		}

		curr = curr.Next
		count++
	}

	//reset pointers
	curr = LL.Head 
	count = func() int{ // adjust count by deleted amount
		if LL.Size % step ==0 {
			return 1
		}

		return LL.Size % step
	}()
	var indexPointer int

	
	for {
		if curr.Next == nil{
			break
		}

		if count % step == 0{
			prev.Next = &nodeArray[len(nodeArray) - 1 - indexPointer]
			indexPointer++
			prev.Next.Next = curr
		}

		prev = *curr
		curr = curr.Next

	}

	return nil
}

func (LN *ListNode) reverse()(*ListNode){
	if LN.Next == nil || LN == nil{
		return LN
	}
	node := LN.Next.reverse()

	node.Next.Next = node

	LN = nil
	
	return node
}

func (LL *LinkedList) Reverse(){
	LL.Head.reverse()
}
