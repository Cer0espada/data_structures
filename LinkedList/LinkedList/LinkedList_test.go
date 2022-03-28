package LinkedList

import (
	"reflect"
	"testing"
)

var testValues []int = []int{11,32,4,5,6,7,8,1,-22,39,1000,52}

func TestAddFront(t *testing.T){
	var LL LinkedList

	for index, test := range testValues{
		LL.AddFront(test)

		if LL.Size !=index +1 {
			t.Errorf("Size is not being incremented correctly, Expected: %v, got %v", index +1, LL.Size)
		}

		if LL.Head.Value != testValues[index]{
			t.Errorf("Head is not being updated as expected, Expected %v, got %v",testValues[index], LL.Head.Value)
		}
	}
}

func TestAddMiddle(t *testing.T){
	var LL LinkedList

	for index, test := range testValues{
		LL.AddMiddle(test)

		node := LL.Head
		for i:= 0; i<LL.Size/2;i++{
			node = node.Next
		}

		if LL.Size !=index +1 {
			t.Errorf("Size is not being incremented correctly, Expected: %v, got %v", index +1, LL.Size)
		}

		if LL.Head.Value != testValues[0]{
			t.Errorf("Head is not being updated as expected, Expected %v, got %v",testValues[0], LL.Head.Value)
		}
	}
}

func TestAddEnd(t *testing.T){
	var LL LinkedList

	for index,test := range testValues{
		LL.AddEnd(test)

		node := LL.Head 

		for i:=0; i<LL.Size-1; i++{
			node = node.Next
		}

		if node.Value != test{
			t.Errorf("The end value doesn't match what is expected, Expected:%v, got %v", test, node.Value,)
		}
		if LL.Size !=index +1 {
			t.Errorf("Size is not being incremented correctly, Expected: %v, got %v", index +1, LL.Size)
		}

		if LL.Head.Value != testValues[0]{
			t.Errorf("Head is not being updated as expected, Expected %v, got %v",testValues[index], LL.Head.Value)
		}
	}
}

func TestRemoveFront(t *testing.T){
	t.Run("testing nil case", func (T *testing.T){
		var LL LinkedList

		err := LL.RemoveFront()

		if err ==nil{
			t.Error("Expected nil case error")
		}
	})

	var LL LinkedList

	for _, test := range testValues{
		LL.AddFront(test)
	}


	for index := range testValues{
		
		if LL.Head.Value != testValues[len(testValues)- 1- index] && len(testValues)- 1- index >1{
			t.Errorf("Head is not being updated appropriately, Expected:%v , got %v", LL.Head.Value, testValues[len(testValues) - 1- index])
		}
		err := LL.RemoveFront()

		if err !=nil{
			t.Error(err)
		}
	
		if LL.Size != len(testValues) - index -1{
			t.Errorf("Size is not being updated as expected, Expected:%v, got %v", LL.Size, len(testValues)-index -1)
		}
	}
}


func TestRemoveMiddle(t *testing.T){
	t.Run("testing nilcase", func(T *testing.T){
		var LL LinkedList

		err:= LL.RemoveMiddle()

		if err == nil{
			t.Error(err)
		}
	})

	var LL LinkedList 

	for _,test := range testValues{
		LL.AddEnd(test)
	}

	for index := range testValues{
		before:= LL.Head
		for i:= 0; i <LL.Size-1; i++{
			before = before.Next
		}
		
		err := LL.RemoveMiddle()

		if err !=nil{
			t.Error(err)
		}
		after:= LL.Head
		for i:= 0; i <LL.Size-1; i++{
			after = after.Next
		}


		if index < len(testValues) -2 && before.Value == after.Value{
			t.Errorf("End is being altered, before and after values match, Before: %v, After: %v",before.Value, after.Value)
		}
		
		if index < len(testValues) -2 && LL.Head.Value != testValues[0]{
			t.Errorf("Expected Head Value to remain unchanged, Expected: %v, got %v", LL.Head.Value, testValues[0])
		} 
	
		if LL.Size != len(testValues) - index -1{
			t.Errorf("Size is not being updated as expected, Expected:%v, got %v", LL.Size, len(testValues)-index -1)
		}
	}
}

func TestRemoveEnd(t *testing.T){
	t.Run("testing nilcase", func(t *testing.T){
				var LL LinkedList

		err:= LL.RemoveEnd()

		if err == nil{
			t.Error(err)
		}
	})

	var LL LinkedList

	for _, test := range testValues{
		LL.AddEnd(test)
	}

	for index := range testValues{

		before:= LL.Head
		for i:= 0; i <LL.Size-1; i++{
			before = before.Next
		}
		
		err := LL.RemoveEnd()

		if err !=nil{
			t.Error(err)
		}
		after:= LL.Head
		for i:= 0; i <LL.Size-1; i++{
			after = after.Next
		}

		if index < len(testValues) -2 && before.Value == after.Value{
			t.Error("End value is not being removed")
		}
		
		if index < len(testValues) -1 && LL.Head.Value != testValues[0]{
			t.Errorf("Head Pointer is being altered, Expected:%v , got %v", testValues[0], LL.Head.Value)
		}

		if LL.Size != len(testValues) - index -1{
			t.Errorf("Size value is not being updated approprately, Expected %v, got %v", len(testValues)-index -1, LL.Size)
		}
		
	}
}

func TestIndexOf(t *testing.T){

	t.Run("testing nilcase", func(t *testing.T){
		var LL LinkedList

		_, err := LL.indexOf(22)

		if err == nil{
			t.Error("Expected nilcase Error")
		}
	})


	var LL LinkedList

	for _, test := range testValues{
		LL.AddEnd(test)
	}

	valx := make([]int,LL.Size)

	for index,test := range testValues{

		val, err := LL.indexOf(test)

		if err != nil{
			t.Error(err)
		}

		valx[val] = test

		if val != index {
			t.Error("Index isn't being appropraitely retrieved")
		}
	}

	if reflect.DeepEqual(valx,testValues) != true{
		t.Error("Retrieved array values don't match")
	}

	val, err := LL.indexOf(200)

	if err ==nil{
		t.Errorf("Expected not found error, instead recieved a Value of %v", val)
	}
}
