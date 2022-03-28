package Trees

import "data_structures/Stacks"

const AlphaSize = 26

type TrieNode struct{
	children [AlphaSize]*TrieNode
	isCompleteWord bool
}

//Trie is a sorted tree-based data-structure that stores the set of strings.
//It has the number of pointers equal to the number of characters of the alphabet in each node. 
//It can search a word in the dictionary with the help of the word's prefix.
type Trie struct{
	Root *TrieNode
	Size int
}

type stringer interface{
	rune 
	string
	int64
}

func (T *Trie) Insert(w string){
	if T.Size == 0{
		T.Root = &TrieNode{children: [AlphaSize]*TrieNode{}, isCompleteWord: false}
		T.Size++
	}
	node := T.Root

	for i:=0;i < len(w);i++{
		charIndex := w[i] - 'a'
		if node.children[int32(charIndex)] == nil{
			node.children[charIndex] = &TrieNode{}
			T.Size++
		}
		node = node.children[charIndex]
		
	}
	node.isCompleteWord = true
}

func (T *Trie) Search(w string) bool{
	node := T.Root

	for i:=0;i < len(w);i++{
		charIndex := w[i] - 'a'
		if node.children[charIndex] == nil{
			return false
		}
		node = node.children[charIndex]
	}
	return node.isCompleteWord
}

func (T *Trie) Delete(w string){
	if len(w) <=0{
		return
	}

	node := T.Root

	 stack := &Stacks.Stack[rune]{Capacity: 100}

	var lastWord TrieNode
	for i:=0;i < len(w);i++{
		charIndex := w[i] - 'a'
		stack.Push(rune(charIndex))
		if node.children[charIndex] == nil{
			return 
		}
		if node.isCompleteWord{
			lastWord = *node
			stack = &Stacks.Stack[rune]{Capacity: 100}
		}
		node = node.children[charIndex]
	}
	node.isCompleteWord = false

	if node.children != [AlphaSize]*TrieNode{}{
		return
	}	

	if node.children == [AlphaSize]*TrieNode{}{
		node = nil
	}
	
	for i:=0;i < stack.Store.Size;i++{
		charIndex, _ := stack.Pop()
		T.Size--
		newNode := lastWord.children[int64(charIndex)]

		if newNode.children == [AlphaSize]*TrieNode{}{
			node = nil
			T.Size = T.Size - stack.Store.Size
			return
		}
	}
}
