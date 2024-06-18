package dsa

/*
** Singly Linked List Node

	It has two prams
	SLLNodeValue: holds the data in the node.
	SLLNextNode: holds the reference of the next node.
*/
type SLLNode[V any] struct {
	SLLNodeValue V
	SLLNextNode  *SLLNode[V]
}
