package Queue

import (
	"data_structures/LinkedList/DoubleLinkedList"
	"testing"
)

var testValues []int = []int{44,33,22,11,7,6,89,3,2,1}

func TestEnqueue(t *testing.T){
	var QQ Queue[int] = Queue[int]{
		Store: DoubleLinkedList.DLinkedList[int]{},
		Capacity: 10,
		Front: 0,
		Rear: 0,
	}

	sq := testValues[:QQ.Capacity]
	for index, test := range testValues{
		err := QQ.Enqueue(test)

		if index +1 < QQ.Capacity{

			if err != nil{
				t.Error(err)
			}
			if QQ.Front != sq[0]{
				t.Errorf("Front is not being intialized properly, Expected %v, got %v",sq[0], QQ.Front)
			}

			if QQ.Rear != sq[index]{
				t.Errorf("Rear is not being intialized properly, Expected %v, bot %v",sq[index], QQ.Rear)
			}
		}
	}
	if !QQ.isFull(){
		t.Errorf("Capacity not being properly maintained")
	}
}

func TestDequeue(t *testing.T){
	var QQ Queue[int] = Queue[int]{
		Store: DoubleLinkedList.DLinkedList[int]{},
		Capacity: 10,
		Front: 0,
		Rear: 0,
	}

	for _, test :=range testValues{
		QQ.Enqueue(test)
	}

	sq := testValues[: QQ.Capacity]

	for index := range testValues{


		value, err := QQ.Dequeue()

		if index +1 > QQ.Capacity{

			if err != nil{
				t.Error(err)
			}
			if value != sq[len(sq) - index - 1]{
				t.Errorf("Return value is not being set properly, Expected %v, got %v",sq[len(sq) - index - 1],value)
			}

			if QQ.Rear != sq[index]{
				t.Errorf("Rear is not being intialized properly, Expected %v, got %v",sq[index], QQ.Rear)
			}

			if index +1 > QQ.Store.Size{
				if QQ.Front != sq[index +2]{
					t.Errorf("Front is not being intialized properly, Expected %v, got %v",sq[len(sq) - index - 1], QQ.Front)
				}
			}
		}
	}

	if !QQ.isEmpty(){
		t.Errorf("Expected Queue to be empty")
	}
}

func TestPeek(t *testing.T){
	var QQ Queue[int] = Queue[int]{
		Store: DoubleLinkedList.DLinkedList[int]{},
		Capacity: 10,
		Front: 0,
		Rear: 0,
	}

	for _, test :=range testValues{
		QQ.Enqueue(test)
	}

	sq:= testValues[:QQ.Capacity]
	
	for index := range sq{

		front, err := QQ.Peek()

		if err != nil{
			t.Error(err)
		}

		if QQ.Rear != sq[len(sq) - 1]{
			t.Errorf("Issue with rear retrieval, Expected: %v, got %v", sq[len(sq) -1], QQ.Rear)
		}
		if front != sq[index]{
			t.Errorf("Issue with rear retrieval, Expected: %v, got %v", sq[len(sq) - index -1], QQ.Rear)
		}

		QQ.Dequeue()
	}


	t.Run("Testing nil case ", func(t *testing.T) {
		var nilcase Queue[int] = Queue[int]{}

		value ,err := nilcase.Peek()

		if err == nil{
			t.Errorf("Failed Nil Case, Expected: %v , got %v", 0, value)
		}
	})

	
}
