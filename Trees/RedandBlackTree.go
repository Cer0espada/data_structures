package Trees

// - Self balancing binary search tree
//
// - The root is black. (Root Property)
// - Every external node is black. (External Property)
// - The children of a red node are black. (Red Property)
// - All external nodes have the same black depth.(Depth Property)
// - Every New node must inserted with Red color.
// - Every leaf (i.e. NULL node) must colored Black.
// - Nodes requireon storage bit to keep track of color
// - The longest path(root farthest nil) is no more than twice the length of the shortest path(root nearest to nil)
// - asymptotic performance
// - Searching in O(log(n))
// - Inserition in O(log(n))
// - Deletion in O(log(n))
type RedAndBlackTreeNode struct{
	Parent 	*RedAndBlackTreeNode
	Left 	*RedAndBlackTreeNode
	Right 	*RedAndBlackTreeNode
	Color string
	Value int 
	Index int 
}

const (
	Black = "Black"
	Red = "Red"
	Left = "Left"
	Right = "Right"
)

func (RBN *RedAndBlackTreeNode) hasParent()(bool){
	if RBN.Parent != nil {
		return true
	}
	return false
}

func (RBN *RedAndBlackTreeNode) hasLeftChild()(bool){
	if RBN.Left != nil {
		return true
	}
	return false
}

func (RBN *RedAndBlackTreeNode) hasRightChild()(bool){
	if RBN.Right != nil {
		return true
	}
	return false
}

func (RBN *RedAndBlackTreeNode) isBlack()(bool){
	if RBN == nil{
		return true
	}
	if RBN.Color == Black{
		return true
	}
	return false
}

func (RBN *RedAndBlackTreeNode) isRed()(bool){
	if RBN ==nil{
		return false
	}

	if RBN.Color == Red{
		return true
	}
	return false
}

func (RBN *RedAndBlackTreeNode) Insert(value int){

	if RBN.Value > value {
		if RBN.hasLeftChild(){
			RBN.Left.Insert(value)
			return
		}
		RBN.Left = &RedAndBlackTreeNode{Parent:RBN, Value:value, Color:Red}

		RBN.Rebalance()
	}

	if RBN.Value < value{
		if RBN.hasRightChild(){
			RBN.Right.Insert(value)
			return
		}
		RBN.Right = &RedAndBlackTreeNode{Parent:RBN, Value:value, Color:Red}
		
		RBN.Rebalance()
	}
}

func (RBN *RedAndBlackTreeNode) Rebalance(){
	
	var direction string

	if RBN.Parent.Left == RBN {
		direction = Left
	}

	if RBN.Parent.Right == RBN{
		direction = Right
	}

	if direction == Left{
		//parent		//child 			//aunt/Uncle
		if RBN.isRed() && RBN.Left.isRed() && RBN.Parent.isBlack() && RBN.Parent.Left == RBN{
		// black line 
		//					a
		//				b	 \
		//			c	 \	  x
		//			 	  y
		//			  
			//Right- Left Rotation
			
			oldLeft := RBN.Left
			oldParent := RBN.Parent
			RBN.Parent = RBN
			RBN.Left = oldParent
			RBN.Left.Right = oldLeft

			//Recolor
			RBN.Color = Black
			RBN.Left.Color = Red 
			RBN.Right.Color = Red
		}
		//newly inserted node has a black aunt/uncle
		//

		if RBN.Parent.Left.isRed(){
			RBN.Color = Black
			RBN.Parent.Left.Color = Black
		}

		if RBN.isRed() && RBN.Left.isRed() && RBN.Parent.Right.isBlack() && RBN.Parent.Right == RBN {
			//black triangle
			//Right - Left Rotation
			oldParent := RBN.Parent
			nodeLeft := RBN.Left
			RBN.Parent = RBN
			RBN.Left = oldParent 
			RBN.Left.Right = nodeLeft
		}
		return 
	}
			//parent		//child 			//aunt/Uncle
			if RBN.isRed() && RBN.Right.isRed() && RBN.Parent.isBlack() && RBN.Parent.Right == RBN {
			// black line 
			//				a
			//			   /   b
			//			  x	  /	 c
			//			 	 y   		
			//			  
				//Left- Right Rotation
				oldRight := RBN.Right
				oldParent := RBN.Parent
				RBN.Parent = RBN
				RBN.Right = oldParent
				RBN.Right.Left = oldRight

				//Recolor
				RBN.Color = Black
				RBN.Left.Color = Red 
				RBN.Right.Color = Red
			}
			//newly inserted node has a black aunt/uncle

			if RBN.Parent.Right.isRed(){
				RBN.Color = Black
				RBN.Parent.Right.Color = Black
			}

			if RBN.isRed() && RBN.Parent.isRed() && RBN.Parent.Left.isBlack() && RBN.Parent.Left == RBN {
				//black triangle
				//Right - Left Rotation
				oldParent := RBN.Parent
				nodeRight:= RBN.Right
				RBN.Parent = RBN
				RBN.Right = oldParent 
				RBN.Right.Left = nodeRight
			}
}

func (RBN *RedAndBlackTreeNode) Delete(value int){

	if RBN == nil {
		return 
	}

	if RBN.Value == value{
		if !RBN.hasLeftChild() || !RBN.hasRightChild(){
			
			RBN = nil
			return 
		}

		if (RBN.hasLeftChild() && !RBN.hasRightChild()) || !RBN.hasLeftChild() && RBN.hasRightChild(){
			if RBN.isBlack() && RBN.Left.isRed(){
				RBN = RBN.Left
				RBN.Color = Black
				return
			}

			if RBN.isBlack() && RBN.Right.isRed(){
				RBN = RBN.Right
				RBN.Color = Black
				return
			}

		}

		if RBN.hasLeftChild() && RBN.hasLeftChild(){
			if RBN.isBlack() && RBN.Parent.Left.isRed(){
				//Left Rotation
				nodeParent := RBN.Parent
				
				oldSibling := RBN.Parent.Left
				RBN.Parent.Left.Left = nodeParent
				nodeParent.Parent = oldSibling
				nodeParent.Right = oldSibling.Left
				

			} 

			if RBN.isBlack() && RBN.Parent.Right.isRed(){
				nodeParent := RBN.Parent
				
				oldSibling := RBN.Parent.Right
				RBN.Parent.Right.Right = nodeParent
				nodeParent.Parent = oldSibling
				nodeParent.Left = oldSibling.Right
			}

			if RBN.isBlack() && RBN.Left.isBlack(){
				// recursively call case 1 on Parent
			}

			if RBN.isBlack() && RBN.Right.isBlack(){
				//recursively call case 1 on Parent
			}

			if RBN.Parent.isRed() && RBN.Parent.Right.isBlack() && RBN.Parent.Right.Right.isBlack() && RBN.Parent.Right.Left.isBlack(){
				RBN.Parent.Color = Black
				RBN.Parent.Right.Color = Red
				return 
			}
			if RBN.Parent.isRed() && RBN.Parent.Left.isBlack() && RBN.Parent.Left.Left.isBlack() && RBN.Parent.Left.Right.isBlack(){
				RBN.Parent.Color = Black
				RBN.Parent.Right.Color = Red
				return 
			}

			if RBN.Parent.Left.isBlack() && RBN.Parent.Left.Right.isRed() && RBN.Parent.Left.Left.isBlack(){
				targetNode := RBN.Parent.Left.Right
				targetNodeLeftChild := RBN.Parent.Left.Right.Left
				targetNodeSibling:= RBN.Parent.Left.Left
				sibling := RBN.Parent.Left
				
				RBN.Parent.Left = targetNode
				
				RBN.Parent.Left.Left = sibling
				RBN.Parent.Left.Right = targetNodeLeftChild
				RBN.Parent.Left.Left.Color = Red
				RBN.Parent.Left.Left.Left = targetNodeSibling
				
				//recoloring
				RBN.Parent.Left.Color = Black
				RBN.Parent.Left.Left.Color = Red
			}

			if RBN.Parent.Right.isBlack() && RBN.Parent.Right.Left.isRed() && RBN.Parent.Left.Right.isBlack(){
				targetNode := RBN.Parent.Right.Left
				targetNodeRightChild := RBN.Parent.Right.Left.Right
				targetNodeSibling:= RBN.Parent.Right.Left
				sibling := RBN.Parent.Right
				
				RBN.Parent.Right = targetNode
				
				RBN.Parent.Right.Right = sibling
				RBN.Parent.Right.Left = targetNodeRightChild
				
				RBN.Parent.Right.Right.Right = targetNodeSibling
				
				//recoloring
				RBN.Parent.Left.Color = Black
				RBN.Parent.Right.Right.Color = Red
			}

			if RBN.Parent.Left.isBlack() && RBN.Parent.Left.Left.isRed() {
				oldParent := RBN.Parent
				sibling := RBN.Parent.Left
				siblingLeftChild := RBN.Parent.Left.Left

				RBN.Parent = sibling 
				RBN.Parent.Right = oldParent
				RBN.Parent.Right.Left = siblingLeftChild

				RBN.Parent.Color = RBN.Parent.Right.Color
			}

			if RBN.Parent.Right.isBlack() && RBN.Parent.Right.Right.isRed() {
				oldParent := RBN.Parent
				sibling := RBN.Parent.Right
				siblingRightChild := RBN.Parent.Right.Right

				RBN.Parent = sibling 
				RBN.Parent.Left = oldParent
				RBN.Parent.Left.Right = siblingRightChild

				RBN.Parent.Color = RBN.Parent.Left.Color
			}
		}
	}

	if RBN.Value > value{
		if RBN.hasLeftChild(){
			RBN.Delete(value)
		}
	}

	if RBN.Value < value{
		if RBN.hasRightChild(){
			RBN.Delete(value)
		}
	}
}

type RedAndBlackTree struct{
	Root *RedAndBlackTreeNode
	Size int
}

func (RBT *RedAndBlackTree) isEmpty()(bool){
	if RBT.Root == nil || RBT.Size ==0 {
		return true
	}
	return false
}

func (RBT *RedAndBlackTree) hasOne()(bool){
	if RBT.Root != nil && RBT.Size ==1 {
		return true
	}
	return false
}
func (RBT *RedAndBlackTree) Insert(value int){
	if RBT.isEmpty(){
		RBT.Root = &RedAndBlackTreeNode{Color:Black,Value:value, Index: 0 }
		RBT.Size ++
		return
	}

	RBT.Root.Insert(value)
}

