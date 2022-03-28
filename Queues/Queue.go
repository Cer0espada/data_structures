package Queue

import (
	"data_structures/LinkedList/DoubleLinkedList"
	"fmt"
)

type queue interface{
	enqueue(int)
	dequeue()(int,error)
	peek()(int,error)
	isFull()(bool, error)
	isEmpty()(bool)

}
//A queue is open at both its ends. One end is always used to insert data (enqueue) and the other is used to remove data (dequeue).
//A Queue follows the First- In-First-Out methodology, ie the data item stored first will be accessed first
//This particular implementation uses a Doubly Linked List which is good for AddFront O(1) and RemoveEnd O(1) Operations
//Enqueue O(1)
//Dequeue O(1)
type Queue[T comparable] struct{
	Store DoubleLinkedList.DLinkedList[T]
	Capacity int
	Front T
	Rear T
}

func (QQ *Queue[T]) isEmpty()(bool){
	if QQ.Store.Head ==nil && QQ.Store.Tail == nil||  QQ.Store.Size == 0{
		return true
	}
	return false
}

func (QQ *Queue[T]) hasOne()(bool){
	if QQ.Store.Head == nil && QQ.Store.Tail == nil || QQ.Store.Size == 1{
		return true
	}
	return false
}

func (QQ *Queue[T]) isFull()(bool){
	if QQ.Store.Size >= QQ.Capacity{
		return true
	}
	return false
}
//Adds elements to the front of the Queue O(1) time
func (QQ *Queue[T]) Enqueue(value T)(error){
	if QQ.isEmpty(){
		QQ.Store.AddFront(value)
		QQ.Front = value
		QQ.Rear = value 
		return nil
	}
	if QQ.isFull(){
		err := fmt.Errorf("Queue is full")
		return err
	}
	QQ.Store.AddFront(value)
	QQ.Rear = value
	return nil
}
//Removes elements to the end of the Queue O(1)
func (QQ *Queue[T]) Dequeue()(T, error){
	if QQ.isEmpty(){
		err:= fmt.Errorf("Queue is currently Empty")
		return *new(T), err
	}

	if QQ.hasOne(){
		value := QQ.Store.Tail.Value
		QQ.Store.RemoveEnd()
		QQ.Front = *new(T)
		QQ.Rear = *new(T)
		return value, nil
	}

	value := QQ.Store.Tail.Value
	QQ.Store.RemoveEnd()
	QQ.Front = QQ.Store.Tail.Value
	return value, nil 
}

//Returns the value of the element end of the queue , returns in O(N) time 
func (QQ *Queue[T]) Peek()(T, error){
	if QQ.isEmpty(){
		err := fmt.Errorf("Queue is empty")
		return *new(T), err
	}

	return QQ.Front, nil
}