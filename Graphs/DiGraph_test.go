package Graphs

import (
	"data_structures/LinkedList/DoubleLinkedList"
	"math/rand"
	"testing"
)

var testValues []int = []int{34,56,2,6,42,3,11,34}
func TestAddNode(t *testing.T) {

	DG := &DGraph{Vertices: make(map[int]DoubleLinkedList.DLinkedList[*Vertex])}

	for index := range testValues{
		DG.AddNode()

		if len(DG.VerticesList) != index +1 && DG.counter != index +1{
			t.Errorf("The Graph Verticies are not being added to the data structure as expected , Expected: %v, got: %v", DG.counter, index+1)
		}
	}
}

func TestRemoveNode(t *testing.T){
	DG := &DGraph{Vertices: make(map[int]DoubleLinkedList.DLinkedList[*Vertex])}

	//generating random not in the graph
	//Adding random edges to test edge removal functionality
	for index := range testValues{
		DG.AddNode()
		DG.AddNode()
		DG.AddEdge(DG.VerticesList[rand.Intn(index+1)], DG.VerticesList[rand.Intn(index+1)])
	}

	for index := range testValues{
		prevVertixLength := len(DG.VerticesList)
		DG.RemoveNode(DG.VerticesList[index])

		if prevVertixLength != len(DG.VerticesList){
			t.Errorf("The Graph Verticies are not being added to the data structure as expected , Expected: %v, got: %v", prevVertixLength, len(DG.VerticesList))
		}
	}

	for index := range DG.VerticesList{
		DG.RemoveNode(DG.VerticesList[index])
	}

	for index := range DG.VerticesList{
		list :=DG.VerticesList[index]

		if list.Adjacent.Size != 0 {
			t.Errorf("Expected list size to be 0, instead got: %v", list.Adjacent.Size)
		}
	}
}
func TestAddEdge(t *testing.T){

	DG := &DGraph{Vertices: make(map[int]DoubleLinkedList.DLinkedList[*Vertex])}

	for index := range testValues{

		v1 := DG.AddNode()
		v2 := DG.AddNode()
		DG.AddEdge(&v1, &v2)

		if len(DG.VerticesList) != 2*index +2 && DG.counter != 2*index +1{
			t.Errorf("The Graph Verticies are not being added to the data structure as expected ")
		}

		if !v1.Adjacent.Contains(&v2){
			t.Errorf("Expected v1 to have v2")
		}

		if !v2.Adjacent.Contains(&v1){
			t.Errorf("Expected v1 to have v2")
		}
	}
}

func TestRemoveEdge(t *testing.T){
	DG := &DGraph{Vertices: make(map[int]DoubleLinkedList.DLinkedList[*Vertex])}

	//generating random not in the graph
	//Adding random edges to test edge removal functionality
	for index := range testValues{
		firstVal := rand.Intn(index+1)
		secondVal := rand.Intn(index +1)
		DG.AddNode()
		DG.AddNode()
		DG.AddEdge(DG.VerticesList[firstVal], DG.VerticesList[secondVal])
		DG.RemoveEdge(DG.VerticesList[firstVal], DG.VerticesList[secondVal])

		if DG.VerticesList[firstVal].Adjacent.Contains(DG.VerticesList[secondVal]){
			t.Errorf("Expected vertex 1 to not have vertex 2")
		}

		if DG.VerticesList[secondVal].Adjacent.Contains(DG.VerticesList[firstVal]){
			t.Errorf("Expected vertex 2 to not have vertex 2")
		}
	}

	for i:= 0; i<len(DG.VerticesList); i++{

		if DG.VerticesList[i].Adjacent.Size != 0{
			t.Errorf("Expected to get size 0 after removal, instead got: %v",  DG.VerticesList[i].Adjacent.Size )
		}
	}

}

func TestBFS(t *testing.T){
	DG := &DGraph{Vertices: make(map[int]DoubleLinkedList.DLinkedList[*Vertex])}

	for i:=0; i<len(testValues);i++{
		DG.AddNode()
		DG.AddNode()
	}

	for index := range DG.VerticesList{
		for i:= 0; i <25;i++{
			DG.AddEdge(DG.VerticesList[rand.Intn(index+1)],DG.VerticesList[rand.Intn(index+1)] )
		}
	}

	for index := range DG.VerticesList{
		
		foundOrigin := DG.BFS(DG.VerticesList[index], DG.VerticesList[len(DG.VerticesList) - 1 - index])

		if foundOrigin <0 {
			t.Errorf("unable to find match , found origin is :%v", foundOrigin)
		}
	}
}

func TestDFS(t *testing.T){
	DG := &DGraph{Vertices: make(map[int]DoubleLinkedList.DLinkedList[*Vertex])}

	for i:=0; i<len(testValues);i++{
		DG.AddNode()
		DG.AddNode()
	}

	// for index := range DG.VerticesList{
	// 	for i:= 0; i <14;i++{
	// 		DG.AddEdge(DG.VerticesList[rand.Intn(index+1)],DG.VerticesList[rand.Intn(index+1)] )
	// 	}
	// }


	DG.AddEdge(DG.VerticesList[0], DG.VerticesList[4])
	DG.AddEdge(DG.VerticesList[4], DG.VerticesList[3])
	DG.AddEdge(DG.VerticesList[3], DG.VerticesList[8])
	DG.AddEdge(DG.VerticesList[8], DG.VerticesList[11])
	DG.AddEdge(DG.VerticesList[11], DG.VerticesList[15])


	v := &Vertex{} // empty node for regular recursive return
	visited := make(map[int]bool)
	DG.DFS(DG.VerticesList[0], DG.VerticesList[15],visited, v)

	if v == nil{
		t.Errorf("Expected DFS to find a node")
	}

	if v != DG.VerticesList[15]{
		t.Errorf("Expected Vertex to be: %v but got :%v", DG.VerticesList[15], v)
	}
}