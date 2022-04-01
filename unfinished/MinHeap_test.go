package Trees

import (
	"testing"
)
var testMHValues []int = []int{1,5,343,7,3,47,69,22}
var sortedMHValues []int = []int{1,3,5,7,22,47,69,343}
var inverseSortedMHValues []int = []int{343,69,47,22,7,5,3,1}
func TestMHInsert(t *testing.T){

	var MHT MinHeapTree

	for index, test := range testMHValues {
		MHT.Insert(test)

		if MHT.Size != index + 1{
			t.Errorf("MHT Size not adjusting as expected, Expected : %v, got: %v", MHT.Size, index+1)
		}

		if MHT.Root.Value != sortedMHValues[0]{
			t.Errorf("MHT Root being altered, Expected:%v, got: %v", sortedMHValues[0], MHT.Root.Value)
		}
	}

}

func TestPeek(t *testing.T){

	t.Run("testing nil case", func(t *testing.T){

		var MHT MinHeapTree

		value, err := MHT.Peek()

		if err == nil{
			t.Error("Expected an error in the nil case")
		}

		if value > 0{
			t.Errorf("Expected value to be: %v instead got %v", -1, value)
		}
	})

	var MHT MinHeapTree

	for _, test := range testMHValues {
		MHT.Insert(test)

		value, err := MHT.Peek()

		if err != nil{
			t.Error(err)
		}

		if MHT.Root.Value != value {
			t.Errorf("Not taking the approprate Value, Expected: %v, got: %v", MHT.Root.Value, value)
		}
	}
}

// func TestPoll(t *testing.T){

// 	t.Run("Testing nil case", func(t *testing.T){
// 		var MHT MinHeapTree

// 		value, err := MHT.Poll()

// 		if err == nil{
// 			t.Errorf("Expected err of %v", err)
// 		}

// 		if value > 0{
// 			t.Errorf("Expected: %v, instead got %v", -1, value)
// 		}
// 	})

// 	t.Run("Test will remain empty", func(t *testing.T){
// 		var MHT MinHeapTree

// 		for _, test := range sortedMHValues{
// 			MHT.Insert(test)
			
// 			value, err := MHT.Poll()

// 			if err != nil{
// 				t.Error(err)
// 			}

// 			if value != test{
// 				t.Errorf("returned minimum is not was expected, Expected: %v but got %v", test, value)
// 			}
			
// 		}
// 		if MHT.Size != 0{
// 			t.Errorf("Error not remaining zero size, Expected: %v, but got %v", 0, MHT.Size)
// 		}
// 	})
	
// 	var MHT MinHeapTree

// 	for _, test := range testMHValues{
// 		MHT.Insert(test)
// 	}

	
// 	for index, test := range sortedMHValues{
// 		value , err := MHT.Poll()
		
// 		if err != nil{
// 			t.Error(err)
// 		}

// 		if test != value {
// 			t.Errorf("min value is not being returned, Expected: %v, got: %v",test, value)
// 		}

// 		if MHT.Size != len(sortedMHValues) - index -1 {
// 			t.Errorf("Size is not being managed as expected, Expected: %v, got: %v", len(sortedMHValues)-index, MHT.Size)
// 		}

// 	}

// }

