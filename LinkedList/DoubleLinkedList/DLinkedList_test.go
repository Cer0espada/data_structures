package DoubleLinkedList

import (
	"fmt"
	"testing"
)

var testValues []int = []int{11,32,3,5,6,7,8,1,-22,4,1000,52}
func TestAddFront(t *testing.T){
	var DLL DLinkedList[int]

	arrayLength := len(testValues) 

	for index, test :=range testValues{
		DLL.AddFront(test)

		if DLL.Size != index + 1 {
			t.Error("Size is not being initialized properly")
		}
		if DLL.Head.Value !=test {
			t.Errorf("Head is not being updated appropriatedly, Expected: %v, got:%v", test, DLL.Head.Value)
		}
	}

	if DLL.Size != arrayLength {
		t.Error("Size not being decremented appropriately")
	}

	if DLL.Head.Value != testValues[arrayLength-1] {
		t.Errorf("Head not being updated appropriately, Expected %v, but got %v", testValues[arrayLength-1], DLL.Head.Value)
	}

	if DLL.Tail.Value != testValues[0] {
		t.Errorf("Tail not being updated appropriately,Expected %v, but got %v", testValues[0], DLL.Head.Value)
	}
}

func TestAddMiddle(t *testing.T){
	//intializing a Head and a Tail
	var DLL DLinkedList[int]
	DLL.AddMiddle(2)
	DLL.AddMiddle(3)

	var offset int = 2

	for _ ,test := range testValues{
		DLL.AddMiddle(test)

		pos ,err := DLL.IndexOf(test)

		if err != nil{
			t.Error(err)
		}


		if pos != handleEven(DLL.Size) {
			t.Errorf("Expected %v but got %v", handleEven(DLL.Size), pos+ 1)
		}
	}

	
	if DLL.Size != len(testValues) + offset {
		t.Error("Size not being incremented appropriately")
	}

	if DLL.Head.Value != 2 || DLL.Tail.Value != 3 {
		t.Error("Tail and Head Value pointers have been modified")
	}
}

func TestAddAtPos(t *testing.T){
	
	posTestTable := struct{
		Input []int
		Pos 	[]int
		Expected []bool
	}{
		Input : testValues,
		Pos : []int{0,0,2,3,2,4,1,2,3,4,5,3},
		Expected :[]bool{true, true, false, true, true, false, true, true, true, true, false},
	}

	var DLL DLinkedList[int]

	for i:= 0 ; i < len(posTestTable.Pos); i++{
		err := DLL.AddAtPos(posTestTable.Pos[i], posTestTable.Input[i])

		if err != nil && !posTestTable.Expected[i] {
			t.Error(err)
		}

		if DLL.Size != i+1{
			t.Errorf("Incorrect Size, Expected: %v but got %v",i+1, DLL.Size )
		}

		
		index, err := DLL.IndexOf(posTestTable.Input[i])

		if err !=nil{
			t.Error(err)
		}

		if index != posTestTable.Pos[i] {
			t.Error("index is not being initialized properly")
		}
	}

	if DLL.Size != len(posTestTable.Input){
		t.Error("LL.Size pointers are not being updated correctly")
	}
}

func TestAddEnd(t *testing.T){
	var DLL DLinkedList[int]
	arrayLength := len(testValues) 

	for _, test := range testValues{
		DLL.AddEnd(test)

		if DLL.Size < 0 {
			t.Error("Size is not being initialized properly")
		}
	}

	if DLL.Size != arrayLength {
		t.Error("Size is not being")
	}

	if DLL.Tail.Value != testValues[len(testValues)-1] {
		t.Errorf("Incorrect tail value, Expected: %v but got %v", testValues[len(testValues)-1], DLL.Tail.Value)
	}
}

func TestIndexOf(t *testing.T){
	 indexOfTestValues := struct{
		input []int
		expected []int
		}{
		input :[]int{2,66,3,7,-22,40,95},
		expected : []int{0,1,2,3,4,5,6},	
	}
	var DLL DLinkedList[int]

	for index, test := range indexOfTestValues.input{
		DLL.AddEnd(test)
	
		pos, err := DLL.IndexOf(test)

		if err !=nil{
			t.Error(err)
		}
		if index != pos{
			t.Errorf("Expected:%v but got: %v", index, pos)
		}
		
	}	
}

func TestRemoveFront(t *testing.T){
	var DLL DLinkedList[int]
	arrayLength := len(testValues) 

	for _, test := range testValues{
		DLL.AddEnd(test)
	}

	for index := range testValues {
		if DLL.Size != arrayLength - index {
			t.Error("Size not being updated appropriately")
		}

		DLL.RemoveFront()

		if index < arrayLength - 2{
			if DLL.Head.Value != testValues[index + 1] {
				t.Errorf("Head not being updated appropriately, Expected:%v, got: %v", testValues[index+1], DLL.Size)
			}	
		}
		
	}
}

func TestRemoveAtPos(t *testing.T){
	
	var DLL DLinkedList[int]

	//initialize LinkedList
	for _, value := range testValues{
		DLL.AddEnd(value)
	}

	for index, test := range testValues{
		node, err := DLL.RemoveAtPos(0)

		if err != nil{
			t.Error(err)
		}

		if node.Value != test{
			t.Errorf("node value does not equal test value, Expected: %v, got: %v ", node.Value, test)
		}

		if DLL.Size != len(testValues) - index - 1{
			t.Errorf("Size is not being initialized properly, Expected: %v, got: %v", len(testValues) - index -1, DLL.Size)
		}

		index, err = DLL.IndexOf(test)

		if err ==nil{
			t.Error(err)
		}

		if index != -1{
			t.Error("value is not being removed as expected")
		}
	}

	if !DLL.isEmpty(){
		t.Error("LL.Size pointer are not being updated correctly")
	}
}

func TestRemoveMiddle(t *testing.T){
	var DLL DLinkedList[int]
	arrayLength := len(testValues) 

	//initialize DLinkedList[int]
	for _, test := range testValues{
		DLL.AddEnd(test)
	}
	
	// maintain that head, tail, and size states are not being mutated 
	for index := range testValues{
		if DLL.Size != arrayLength - index {
			t.Error("Size not being updated appropriately")
		}

		DLL.RemoveMiddle()

		if arrayLength - index >= 2 {
			if DLL.Head.Value != testValues[0] || DLL.Tail.Value != testValues[len(testValues)-1] {
				t.Error("Head and Tails states are being updated erroneously")
			}
		}

		if DLL.Size != arrayLength - index -1 {
			t.Errorf("Size state is not being properly maintained, Expected: %v , got: %v", arrayLength - index -1, DLL.Size)
		}
	}

	//maintain that the middle is indeed being Removed
	
	curr := DLL.Head

	for i :=0; i< DLL.Size / 2; i++{
			curr = curr.Next
		}
		testValue := curr.Value

		DLL.RemoveMiddle()

		curr = DLL.Head
		fmt.Println(curr)

		for{
			if curr.Next == nil || curr.Value == testValue {
				break
			}
			curr = curr.Next
		}
}
func TestRemoveEnd(t *testing.T){
	var DLL DLinkedList[int]
		

		for _, test := range testValues{
			DLL.AddEnd(test)
		}

		for index := range testValues{

			var expectedTailIndex = len(testValues) - index -2 
			var expectedDLLSize = len(testValues) - index 
	
			
			if DLL.Size != expectedDLLSize {
				t.Errorf("Size state is not being properly maintained, Expected: %v, got: %v", expectedDLLSize, DLL.Size)
			}
			err := DLL.RemoveEnd()

			if err != nil{
				t.Error(err)
			}

			if  DLL.Size > 2 && DLL.Tail.Value != testValues[expectedTailIndex] {
				t.Errorf("Tail is being updated unappropriately, Expected: %v, got: %v", testValues[expectedTailIndex], DLL.Tail.Value)
			}
		}
}

func handleEven(ListSize int) (int) {
		if ListSize % 2 == 0 {
			return ListSize /2 -1
		}
		return ListSize / 2
}