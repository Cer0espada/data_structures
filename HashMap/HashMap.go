package hashmap

import (
	"data_structures/LinkedList/DoubleLinkedList"
	"fmt"
)

type keyer interface{
	comparable
	string  
}
const storeSize = 4

type accessPair[K keyer, V comparable] struct{
	key   *DoubleLinkedList.DLinkedList[K]
	value *DoubleLinkedList.DLinkedList[V]
}

type HashMap[K keyer, V comparable] struct{
	Store [storeSize]*accessPair[K,V]
	Size int 
}

func (HM *HashMap[K,V]) isEmpty()(bool){
	if HM.Size == 0 {
		return true
	}
	return false
}

//this simple hash function returns the remainder of the converted byte array into int64
func hash[K keyer](word K)(int){
	wordBytes := []byte(string(word))

	total := 0 
	for _, val := range wordBytes{
		total = int(val) + total
	}
	return total % storeSize
}

func (HM *HashMap[K,V]) Add(key K , value V){

	
	if HM.Store[hash(key)]==nil{

		newKeyNode := &DoubleLinkedList.DListNode[K]{Value:key}
		newValueNode := &DoubleLinkedList.DListNode[V]{Value:value}
		HM.Store[hash(key)] = &accessPair[K,V]{
				key: &DoubleLinkedList.DLinkedList[K]{
					Head:newKeyNode,
					Tail:newKeyNode,
					Size:1,
				},
				value: &DoubleLinkedList.DLinkedList[V]{
					Head:newValueNode,
					Tail:newValueNode,
					Size:1,
				},
		}
		HM.Size++
		return
	}

	keyValue := HM.Store[hash(key)]

	keyValue.key.AddEnd(key)
	keyValue.value.AddEnd(value)
	HM.Size++
}

func (HM *HashMap[K,V]) Search(key K)(bool){
	if HM.isEmpty(){
		return false
	}

	keyValue := HM.Store[hash(key)]

	if keyValue == nil {
		return false
	}

	if keyValue.key.Contains(key){
		return true
	}

	return false
}

func (HM *HashMap[K,V]) Get(key K) (V,error){
	if HM.isEmpty(){
		err:= fmt.Errorf("hashmap is empty")
		return *new(V), err 
	}
	
	keyValue := HM.Store[hash(key)] // linkedlist index 

	keyNode, err := keyValue.key.Find(key) // 

	if err != nil{
		return *new(V), err
	}
	
	KeyIndex, err := keyValue.key.IndexOf(keyNode.Value)
	if err != nil{
		return *new(V), err
	}
	ValueNode, err := keyValue.value.GetIndex(KeyIndex)

	if err != nil{
		return *new(V), err
	}
	valueIndex, err := keyValue.value.IndexOf(ValueNode.Value)
	if KeyIndex == valueIndex{
		return ValueNode.Value, nil 
	}
	err = fmt.Errorf("Indexes dont match ")
	return *new(V), err
}

func (HM *HashMap[K,V]) Delete(key K)(error){
	if HM.isEmpty(){
		err := fmt.Errorf("hashmap is emtpy")
		return err
	}

	keyValue := HM.Store[hash(key)] // linkedlist index 

	keyNode, err := keyValue.key.Find(key)

	if err != nil{
		return err
	}
	
	keyIndex, err := keyValue.key.IndexOf(keyNode.Value)

	if err != nil{
		return  err
	}
	ValueNode, err := keyValue.value.GetIndex(keyIndex)

	if err != nil{
		return  err
	}
	valueIndex, err := keyValue.value.IndexOf(ValueNode.Value)

	if err != nil{
		return err 
	}

	if keyIndex != valueIndex{
		err := fmt.Errorf("Retrieved unexpected value")
		return err
	}
	

	keyValue.key.RemoveAtPos(keyIndex)
	keyValue.value.RemoveAtPos(valueIndex)

	HM.Size--

	return nil
}