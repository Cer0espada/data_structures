package Stacks

import (
	"data_structures/LinkedList/DoubleLinkedList"
	"fmt"
)

type stack interface{
	push(int)
	pop(int)  (int, error)
	isFull()  (bool)
	isEmpty() (bool)
	peek() (int)
}

//A Stack is a collection data structure in which only the top or first item can be accessed at a time
//It is a First-In-First-Out Data Stucture 
//This implementation of a stack uses a Doubly LinkedList for easy addFront O(1) and removeFront O(1) Operation
type Stack[T comparable] struct{
	Store DoubleLinkedList.DLinkedList[T]
	Capacity int 
	Top T
}

func (SS *Stack[T]) isEmpty()(bool){
	if SS.Store.Size == 0 || SS.Store.Head == nil && SS.Store.Tail == nil{
		return true
	}
	return false
}

func (SS *Stack[T]) hasOne()(bool){
	if SS.Store.Size == 1 && SS.Store.Head == SS.Store.Tail{
		return true
	}
	return false
}

func (SS *Stack[T]) isFull()(bool){
	if SS.Store.Size  >= SS.Capacity{
		return true
	}
	return false
}
//Push O(1) Puts element on the top of the stack
func (SS *Stack[T]) Push(value T)(error){

	if SS.Store.Size < SS.Capacity{
		SS.Store.AddFront(value)
		SS.Top = value
		return nil
	}
	err := fmt.Errorf("Stack is full, Capacity is currently %v, there are %v items in the Stack", SS.Capacity, SS.Store.Size)
	return err
}

//Pop O(1) Takes the element from the top of the stack and removes it and returns that element
func (SS *Stack[T]) Pop() (T,error){

	if SS.isEmpty(){
		err :=fmt.Errorf("Stack size is empty")
		return *new(T), err
	}

	if SS.Store.Size ==1 {
		top := SS.Store.Head.Value
		SS.Store.RemoveFront()
		return top, nil
	}
	top := SS.Store.Head.Value
	SS.Store.RemoveFront()
	SS.Top = SS.Store.Head.Value
	return top, nil
}

//Peek O(1) Takes the element from the top of the stack and returns it
func (SS *Stack[T]) Peek() (T, error){
	if SS.Store.Head == nil{
		err := fmt.Errorf("Stack is empty")
		return *new(T), err
	}

	return SS.Top, nil 
}
