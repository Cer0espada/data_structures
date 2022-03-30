package DoubleLinkedList

import (
	"errors"
	"fmt"
)

//ListNode with Prev Node Optimization
type DListNode[T comparable] struct {
	Prev  *DListNode[T]
	Next  *DListNode[T]
	Value T
}

//Double linked lists are an optimization over the traditional LinkedList Data type
//There is a bidirectional track that we can traverse the list in; we start at the head node and move to a tail node
//traverse from the root until the last node, which will end at an empty null value.
//Includes a Tail Optimization to ge the last element , converting an O(N) to a O(1) Lookup
//Linked List tend to be more memory efficent and sacrifice the typical O(1) index look up that Arrays/ Lists /Slices have 
type DLinkedList[T comparable] struct {
	Head *DListNode[T]
	Tail *DListNode[T]
	Size int
}

func (DLL *DLinkedList[T]) isEmpty()(bool){
	if DLL.Head == nil && DLL.Tail == nil || DLL.Size ==0{
		return true
	}
	return false
}

func (DLL *DLinkedList[T]) hasOne()(bool){
	if DLL.Head != nil && DLL.Head == DLL.Tail|| DLL.Size ==1{
		return true
	}
	return false
}
//Insert at Front O(1)
func (DLL *DLinkedList[T]) AddFront(Value T) {
	newNode := DListNode[T]{nil, nil, Value}
	if DLL.isEmpty(){
		DLL.Head = &newNode
		DLL.Tail = &newNode
		DLL.Size++
		return 
	}
	DLL.Head.Prev = &newNode
	newNode.Next = DLL.Head
	DLL.Head = &newNode
	DLL.Size++
}
//Insert at N position with value of T
//
//Time Complexity  O(N)
func (DLL *DLinkedList[T]) AddAtPos(pos int, value T)(error){
	if DLL.isEmpty(){
		DLL.AddFront(value)
		err := fmt.Errorf("warn: LinkedList is empty")
		return err
	}

	if pos >= DLL.Size {
		err:= fmt.Errorf("index position %v is greater than Linked List's current size: %v", pos, DLL.Size)
		return err
	}

	if pos == 0{
		DLL.AddFront(value)
		return nil
	}

	if pos == DLL.Size - 1{
		DLL.AddEnd(value)
		return nil
	}

	if pos == DLL.Size / 2 {
		DLL.AddMiddle(value)
		return nil
	}
	curr := DLL.Head

	for i :=0; i<pos-1; i++ {
		if curr.Next == nil{
			break
		}
		curr = curr.Next
	}

	newNode := &DListNode[T]{Value:value}
	nextNode := curr.Next

	curr.Next = newNode
	newNode.Prev = curr
	newNode.Next = nextNode
	nextNode.Prev = newNode
	DLL.Size++
	return nil
}
//Insert at end O(1)
func (DLL *DLinkedList[T]) AddEnd(Value T) {
	
	if DLL.isEmpty() {
		newNode := DListNode[T]{nil, nil, Value}
		DLL.Head = &newNode
		DLL.Tail = &newNode
		DLL.Size++
		return
	}

	if DLL.hasOne(){
		newNode := DListNode[T]{DLL.Head, nil, Value}
		DLL.Head.Next = &newNode
		DLL.Tail = &newNode
		DLL.Size++
		return 
	}
	node := &DListNode[T]{Value:Value}

	DLL.Tail.Next = node
	node.Prev = DLL.Tail.Next
	DLL.Tail = node 
	DLL.Size++
}

//Unstable insertion of value into the middle of the Linked List
func (DLL *DLinkedList[T]) AddMiddle(Value T) {
	
	if DLL.isEmpty() {
		DLL.AddFront(Value)
		return 
	}

	if DLL.hasOne() {
		DLL.AddEnd(Value)
		return
	}

	curr := DLL.Head

	mid := DLL.Size /2 

	var count int = 1
	for count < mid {
		curr = curr.Next
		count++
	}
	nextNode := curr.Next
	newNode := DListNode[T]{curr, curr.Next, Value}

	nextNode.Prev = &newNode
	newNode.Prev = curr
	curr.Next = &newNode
	DLL.Size++ 
}
// retrieves the index or position of a specific node value in the LinkedList
//
//Time Complexity is O(N)
func (DLL *DLinkedList[T]) IndexOf(Value T) (int, error) {
	if DLL.isEmpty() {
		err := fmt.Errorf("linkedlist is empty at this position")
		return -1, err
	}

	if DLL.hasOne(){
		if DLL.Head.Value != Value{
			err := fmt.Errorf("Value does not exist in this List")
			return -1, err
		}

		return 0, nil
	}

	var count int
	curr := DLL.Head

	for curr.Next != nil{

		if curr.Value == Value {
			return count, nil
		}

		if curr.Next == nil {
			break
		}
		curr = curr.Next
		count++
	}

	if curr.Value != Value{
		err:= fmt.Errorf("code should be unreachable at this point, count: %v", count)
		return -1, err
	}

	return count, nil

	
}
//Retrieves the node given a certain index
//
//Time complexity O(N)
func (DLL *DLinkedList[T]) GetIndex(index int) (*DListNode[T], error){
	if DLL.isEmpty(){
		err := fmt.Errorf("LinkedList is empty at this position")
		return nil, err
	}

	if index < 0 || index > DLL.Size-1{
		err := fmt.Errorf("Not a valid index number")
		return nil, err
	}

	count := index
	curr := DLL.Head

	for count >0{
		curr = curr.Next
		count--
	} 

	return curr,nil
}

func (DLL *DLinkedList[T]) PrintList() {
	if DLL.Head == nil {
		fmt.Println("Looks like notihing is here")
	}

	curr := DLL.Head

	for curr.Next != nil{
		fmt.Println(curr.Value)
		curr = curr.Next
	}
	fmt.Println("thats the whole list")
}

//Remove at Front, removes current head node
//
//Time Complexity is O(1)
func (DLL *DLinkedList[T]) RemoveFront() (error){
	if DLL.isEmpty() {
		err := fmt.Errorf("Linked List is empty")
		return err
	}

	if DLL.hasOne() {
		DLL.Tail = nil
		DLL.Head = nil
		DLL.Size--
		return nil
	}

	curr := DLL.Head.Next

	DLL.Head= nil
	curr.Prev = nil
	DLL.Head = curr
	DLL.Size--
	return nil
}
//Remove at End 
//
//Time Complexity is O(1)
func (DLL *DLinkedList[T]) RemoveEnd()(error) {
	if DLL.isEmpty() {
		err := fmt.Errorf("linkedList is empty")
		return err
	}

	if DLL.hasOne() {
		DLL.Head = nil 
		DLL.Tail = nil 
		DLL.Size--
		return nil
	}

	lastNode := DLL.Tail.Prev
	DLL.Tail.Prev = nil
	lastNode.Next = nil 
	DLL.Tail = lastNode
	DLL.Size--
	return nil
}

//unstable removal of middle node of children
func (DLL *DLinkedList[T]) RemoveMiddle()(error) {
	if DLL.isEmpty() {
		err := errors.New("The LinkedList is empty")
		return err
	}

	if DLL.hasOne() {
		temp := DLL.Tail.Prev
		temp.Next = nil
		DLL.Tail.Prev = nil
		DLL.Tail = temp
		DLL.Size--
		return nil
	}

	curr := DLL.Head

	mid := DLL.Size / 2

	for i := 0; i >= mid; i++ {
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}

	curr.Next.Prev = curr.Prev
	curr.Prev = curr.Next
	DLL.Size--
	return nil
}
//Removal of value at Position N 
//
//Time complexity O(N)
func (DLL *DLinkedList[T]) RemoveAtPos(index int)(*DListNode[T], error){
	if DLL.isEmpty(){
		err:= fmt.Errorf("Linked List is empty")
		return nil,err
	}

	if index < 0{
		err := fmt.Errorf("Indicated string is not present in the index")
		return nil, err
	}

	if index > DLL.Size -1{
		err := fmt.Errorf("Indicated string exceeds expected index range")
		return nil, err
	}

	if index ==0{
		temp := DLL.Head
		DLL.RemoveFront()
		return temp, nil
	}

	if index == DLL.Size -1 {
		temp := DLL.Tail 
		DLL.RemoveEnd()
		return temp, nil
	}

	curr := DLL.Head
	for i := 0; i<index;i++{
		curr = curr.Next
	}

	prev := curr.Prev
	nextNode := curr.Next

	prev.Next = nextNode
	nextNode.Prev = prev

	DLL.Size--
	return curr, nil
}

//Finds and returns value in LinkedList
func (DLL *DLinkedList[T]) Find(value T)(*DListNode[T], error){
	if DLL.isEmpty(){

		err:= fmt.Errorf("Linked List is empty")
		return nil, err
	}

	if DLL.hasOne(){
		if DLL.Head.Value != value {
			err := fmt.Errorf("Value not present in LinkedList")
			return nil, err 
		}

		return DLL.Head, nil
	}

	curr := DLL.Head 

	for curr.Next != nil{
		if curr.Value == value{
			return curr, nil
		}

		curr = curr.Next
	}

	if curr.Value !=value{
		err := fmt.Errorf("Value of: %v does not exist in Linked List", value)
		return curr, err 
	}

	
	return curr, nil
}


func (DLL *DLinkedList[T]) DetectCycle() (bool, error) {
	if DLL.isEmpty() {
		fmt.Println("DLinkedList[T] is empty")
		return false, nil
	}

	if DLL.hasOne() {
		return false, nil
	}

	// checks if the next pointer and the prev pointer alignment
	var PrevNodeArray []T = make([]T, 0, DLL.Size)
	var NextNodeArray []T = make([]T, 0, DLL.Size)

	fpcurr := DLL.Head // forward pointer
	epcurr := DLL.Tail // end pointer

	for i:= DLL.Size; i <0;i--{

		if fpcurr.Next == nil || epcurr ==nil{
			err:= fmt.Errorf("next and Prev pointers are not being initialized correctly")
			return false, err
		}

		PrevNodeArray = append(PrevNodeArray, epcurr.Value)
		NextNodeArray = append(NextNodeArray, fpcurr.Value)

		fpcurr  = fpcurr.Next
		epcurr = epcurr.Prev

	}

	if len(PrevNodeArray) != len(NextNodeArray) {
		return true, nil
	}


	fast := DLL.Head
	slow := DLL.Head

	for i := 0; i < DLL.Size; i++ {

		if slow.Next == nil || fast.Next.Next == nil {
			fmt.Println("No cycle detected")
			return false, nil
		}

		if slow == fast || fast.Next == fast.Prev {
			return true, nil
		}

		slow = slow.Next
		fast = fast.Next.Next

	}

	if slow != DLL.Tail || fast == DLL.Tail {
		err := errors.New("loop pointers aren't aligned")
		return false, err
	}

	return false, nil
}
//returns bool value whether value is present in the LinkedList
//
//Time Complexity O(N)
func (DLL *DLinkedList[T]) Contains(value T)(bool){

	if DLL.isEmpty(){
		
		return false 
	}

	if DLL.hasOne(){
		if DLL.Head.Value == value {
			return true
		}
		return false
	}

	curr := DLL.Head 

	for curr.Next != nil{
		if curr.Value == value{
			return true
		}
		
		curr = curr.Next
	}

	if curr.Value != value {
		return false
	}

	return true
}

func (DLL *DLinkedList[T]) ToSlice()([]T, error){
	if DLL.isEmpty(){
		err := fmt.Errorf("LinkedList is empty")
		return nil, err
	}

	if DLL.hasOne(){
		return []T{DLL.Head.Value}, nil
	}
	curr := DLL.Head

	DLLX := make([]T,0,DLL.Size)

	for{
		if curr.Next == nil{
			break
		}
		DLLX = append(DLLX, curr.Value)
		curr = curr.Next
	}
	return DLLX, nil
}

