package Stacks

import (
	"data_structures/LinkedList/DoubleLinkedList"
	"testing"
)

var testValues []int = []int{45,6,232,5,9,5,0,1,2,4,34}

func TestPush(t *testing.T){

	var SS Stack[int] = Stack[int]{
		Store: DoubleLinkedList.DLinkedList[int]{nil, nil,0},
		Capacity: 10,
	}

	for index, test := range testValues{

		err := SS.Push(test)

		if err != nil && SS.Capacity >= index +1{
			t.Error(err)
		}

		if SS.Store.Size != index + 1 && index+1 < SS.Capacity {
			t.Errorf("Size is not being updated correctly, Expected:%v, got %v", SS.Store.Size, index+1)
		}

		if SS.Store.Size > SS.Capacity{
			t.Errorf("Size is not being updated correctly, Expected:%v ,got %v", SS.Store.Size, SS.Capacity)
		}
	}

	cs := testValues[:SS.Store.Size]

	if SS.Store.Tail.Value != cs[0]{
		t.Errorf("Tail Value has gone beyond stack capacity, Expected %v, got %v",cs[0], SS.Store.Tail.Value)
	}
}

func TestPop(t *testing.T){
	var SS Stack[int] = Stack[int]{
		Store: DoubleLinkedList.DLinkedList[int]{nil, nil,0},
		Capacity: 10,
	}


	for _, test :=range testValues{
		SS.Push(test)
	}
	cs := testValues[:SS.Store.Size]

	for index := range testValues{
		value, _ := SS.Pop()

		if SS.Capacity > index + 1{
			if testValues[len(cs)-1 -index] != value {
				t.Errorf("Values not being popped in correct order, Expected %v, got %v",testValues[len(cs)-1 -index], value )
			}

		}

		
	}

	if !SS.isEmpty(){
		t.Errorf("Size and Item Quantity is not being intialized correctly, Expected: %v as Size, got %v as Size",0,SS.Store.Size)
	}
	
}

func TestPeek(t *testing.T){
	var SS Stack[int] = Stack[int]{
		Store: DoubleLinkedList.DLinkedList[int]{nil, nil,0},
		Capacity: 10,
	}


	for _, test :=range testValues{
		SS.Push(test)
	}

	cs:= testValues[:SS.Capacity]
	for index := range cs{

		top, err := SS.Peek()

		if err != nil{
			t.Error(err)
		}

		if top != cs[len(cs) - index - 1]{
			t.Errorf("Issue with Top retrival, Expected: %v, got %v", cs[len(cs) - index -1], top)
		}

		SS.Pop()
	}


	t.Run("Testing nil case ", func(t *testing.T) {
		var nilcase Stack[int] = Stack[int]{}

		value ,err := nilcase.Peek()

		if err == nil{
			t.Errorf("Failed Nil Case, Expected: %v , got %v", 0, value)
		}
	})

}