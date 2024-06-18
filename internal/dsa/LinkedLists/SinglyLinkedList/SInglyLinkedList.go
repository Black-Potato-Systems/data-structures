/*
**

	This file contains the implmentation of singly linked list.
*/
package singlylinkedlist

import (
	"reflect"

	"github.com/blackpotato/data-structures/internal/dsa"
	"github.com/golang/glog"
)

// SLLStruct defines singly linked list struct
type SLLStruct[V any] struct {
	Head *dsa.SLLNode[V]
	Tail *dsa.SLLNode[V]
	Size int
}

// SLLInterface contains all SLL methods
type SLLInterface[V any] interface {
	CreateSLL(element V)
	SLLHead() *dsa.SLLNode[V]
	SLLTail() *dsa.SLLNode[V]
	SLLSize() int
	Insert(element V, location int)
	Delete(location int)
	Traverse()
	Search(element V) bool
}

// InitSLL initializes Singly Linked List
func InitSLL[V any]() SLLInterface[V] {
	return &SLLStruct[V]{}
}

// CreateSLL create SLL with the provided element
func (sll *SLLStruct[V]) CreateSLL(element V) {

	node := &dsa.SLLNode[V]{
		SLLNodeValue: element,
		SLLNextNode:  nil,
	}

	sll.Head = node
	sll.Tail = node
	sll.Size = 1
}

// function returns the head of the linked list.
func (sll *SLLStruct[V]) SLLHead() *dsa.SLLNode[V] {
	return sll.Head
}

// function returns the tail of the linked list
func (sll *SLLStruct[V]) SLLTail() *dsa.SLLNode[V] {
	return sll.Tail
}

// function returns the size of the linked list
func (sll *SLLStruct[V]) SLLSize() int {
	return sll.Size
}

// insert an element at the given element of the linked list
func (sll *SLLStruct[V]) Insert(element V, location int) {

	if sll.Head == nil {
		sll.CreateSLL(element)
	}

	// inserting an element at the start of the linked list
	if location == 0 {
		node := &dsa.SLLNode[V]{
			SLLNodeValue: element,
			SLLNextNode:  sll.Head,
		}
		sll.Head = node
		sll.Size++
	} else if location >= sll.Size {
		node := &dsa.SLLNode[V]{
			SLLNodeValue: element,
			SLLNextNode:  nil,
		}
		sll.Tail.SLLNextNode = node
		sll.Size++
	} else {

		tempNode := sll.Head
		for i := 0; i < location; i++ {
			tempNode = tempNode.SLLNextNode
		}

		node := &dsa.SLLNode[V]{
			SLLNodeValue: element,
			SLLNextNode:  tempNode.SLLNextNode,
		}
		tempNode.SLLNextNode = node
		sll.Size++
	}
}

// function deletes the element at the provided location.
func (sll *SLLStruct[V]) Delete(location int) {

	if sll.Head == nil {
		glog.Infoln("empty list")
	}

	if location == 0 {
		sll.Head = sll.Head.SLLNextNode
		sll.Size--

		if sll.Size == 0 {
			sll.Head = nil
			sll.Tail = nil
			sll.Size = 0
		}
	} else if location >= sll.Size {

		tempNode := sll.Head
		for i := 0; i < sll.Size-1; i++ {
			tempNode = tempNode.SLLNextNode
		}
		tempNode.SLLNextNode = nil
		sll.Tail = tempNode
		sll.Size--

		if sll.Size == 0 {
			sll.Head = nil
			sll.Tail = nil
			sll.Size = 0
		}

	} else {
		tempNode := sll.Head
		for i := 0; i < location; i++ {
			tempNode = tempNode.SLLNextNode
		}
		tempNode.SLLNextNode = tempNode.SLLNextNode.SLLNextNode
		sll.Size--
	}

}

// function prints all the elemetns of the linked list.
func (sll *SLLStruct[V]) Traverse() {
	tempNode := sll.Head
	for i := 0; i < sll.Size; i++ {
		glog.Info(tempNode.SLLNodeValue)
		tempNode = tempNode.SLLNextNode
	}
}

// function searches an element in the linked list
func (sll *SLLStruct[V]) Search(element V) bool {

	var isFound bool

	tempNode := sll.Head

	for i := 0; i < sll.Size; i++ {
		if reflect.DeepEqual(tempNode.SLLNodeValue, element) {
			isFound = true
			return isFound
		} else {
			tempNode = tempNode.SLLNextNode
		}
	}
	return isFound
}
