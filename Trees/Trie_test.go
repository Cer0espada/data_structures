package Trees

import (
	"fmt"
	"testing"
)

//+4		//+3 		//+3 	//+3
var testWordValues []string = []string{"race", "racecar", "rose", "rosette"}

func TestTrieInsert(t *testing.T){
	
	var TT Trie

	for _, test :=range testWordValues{
		TT.Insert(test)
	}

	for _, test :=range testWordValues{
		node := TT.Root

		for i:=0; i< len(test); i++{
			charIndex := test[i] - 'a'

			if node.children[charIndex] == nil && node.isCompleteWord==false{
				err := fmt.Errorf("Expected character to be present in structure")
				t.Error(err)
			}

			node = node.children[charIndex]
		}
	}

	if TT.Size != 14{
		t.Errorf("Size is not being initialized correctly, Expected:%v got:%v", 14, TT.Size)
	}

	if len(TT.Root.children) != 26{
		t.Errorf("Values not being inserted as Expected, Expected: %v, got: %v", 1, len(TT.Root.children))
	}

}

func TestTrieSearch(t *testing.T){
	var TT Trie

	for _, test :=range testWordValues{
		TT.Insert(test)
	}

	for _,test := range testWordValues{
		found := TT.Search(test)

		if !found{
			t.Errorf("Expected %v to be in the Trie", test)
		}
	}
}

