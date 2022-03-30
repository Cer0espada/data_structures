package Graphs

import (
	"data_structures/LinkedList/DoubleLinkedList"
	Queue "data_structures/Queues"
)

type Vertex struct{
	Adjacent DoubleLinkedList.DLinkedList[*Vertex]
	Weight 	int 
	Origin 	int
}


type DGraph struct{
	Vertices map[int]DoubleLinkedList.DLinkedList[*Vertex]
	Weight 	int 
	counter int 
	VerticesList []*Vertex
}
//Adds a Vertex to the graph stuct without any connections / edges
func (DG *DGraph) AddNode()(Vertex){
	node :=&Vertex{Origin: DG.counter}
	DG.VerticesList = append(DG.VerticesList, node)
	DG.counter++
	return *node
}

func (DG *DGraph) RemoveNode(v *Vertex){

	iter := v.Adjacent.Head

	if iter == nil{
		return
	}
	//iterates through each v reference in the adjacenylist removes it 
	count := 0
	for iter.Next != nil {
		pos, _ := (iter.Value).Adjacent.IndexOf(v)
		iter.Value.Adjacent.RemoveAtPos(pos)
		iter = iter.Next
		count++
	}
	//delete value from map origin
	DG.VerticesList[v.Origin] = &Vertex{}
}

func (DG *DGraph) AddEdge(v1 *Vertex, v2 *Vertex){

	if v2.Adjacent.Contains(v1) || v1.Adjacent.Contains(v2){
		return
	}
	v2.Adjacent.AddEnd(v1)
	v1.Adjacent.AddEnd(v2)
	
	graphV1 := DG.Vertices[v1.Origin]
	graphV1.AddEnd(v2)
	DG.Vertices[v1.Origin] = graphV1

	graphV2 := DG.Vertices[v2.Origin]
	graphV2.AddEnd(v1)
	DG.Vertices[v2.Origin] = graphV2
}

func (DG *DGraph) RemoveEdge(v1 *Vertex, v2 *Vertex){

	//position of v1 in v2's linked list and vice versa
	//Remove reference of the edge on the internal nodes
	posV1, _ := v2.Adjacent.IndexOf(v1)
	v2.Adjacent.RemoveAtPos(posV1)
	posV2, _ := v1.Adjacent.IndexOf(v2)
	v1.Adjacent.RemoveAtPos(posV2)

	//Remove reference of the edge from main graph parent v2 to v1
	OriginV2 := DG.Vertices[v2.Origin]
	valV1 ,_ := OriginV2.IndexOf(v1)
	OriginV2.RemoveAtPos(valV1)

	//Remove reference of the edge from main graph parent v1 to v2
	OriginV1 := DG.Vertices[v1.Origin]
	valV2 ,_ := OriginV2.IndexOf(v2)
	OriginV1.RemoveAtPos(valV2)

	//reset the references on the map object
	DG.Vertices[v2.Origin]= OriginV1
	DG.Vertices[v1.Origin] = OriginV2
}


//Breath First Search
//
//Time complexity O(V + E) (Vertices and Edges)
//BFS is an algorithm for graph traversal in which you are tracing the amount of paths to get from point A to point B
func (DG *DGraph) BFS(node *Vertex, Value *Vertex)(int) {

	visited := make(map[int]bool) // map that acts as a set

	queue:= &Queue.Queue[*Vertex]{Capacity: 1000}
	queue.Enqueue(node)

	for queue.Store.Size > 0 {
		V, _ := queue.Dequeue()
		list := DG.Vertices[V.Origin]
		
		head := list.Head
		for head.Next != nil {

			head = head.Next

			if head.Value == Value{
				return head.Value.Origin
			}

			if visited[head.Value.Origin] != true{
				visited[head.Value.Origin] = true
				queue.Enqueue(head.Value)
			}
		}
	}
	return -1
}

func initVisited(v map[int]bool)(map[int]bool){
	if v == nil{
		v = make(map[int]bool)
		return v
	}
	return v
}

//Depth First Search
//
//Time complexity O(V + E) (Vertices and Edges)
//BFS is an algorithm for graph traversal to find the shortest path from point A to point B
func (DG *DGraph) DFS(node *Vertex, Value *Vertex, v map[int]bool, foundNode *Vertex){
	visited := initVisited(v)

	if visited[node.Origin] != true{
		visited[node.Origin] = true
	}

	list := DG.Vertices[node.Origin]


	head := list.Head

	for head != nil{
		if head.Value == Value{
			*foundNode = *Value 
			return
		}

		if !visited[head.Value.Origin]{
			DG.DFS(head.Value, Value, visited, foundNode)
		}
		head = head.Next
	}
}

func (DG *DGraph) getGraphSize ()(int){
	return len(DG.VerticesList)
}