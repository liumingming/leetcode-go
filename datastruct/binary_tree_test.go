package datastruct

import "testing"

func TestNewBinaryTree_MidOrder(t *testing.T) {
	tree := NewBinaryTree()
	tree.AddNode(3)
	tree.AddNode(1)
	tree.AddNode(2)
	tree.AddNode(5)
	tree.AddNode(4)
	//
	tree.MidOrder()
}


func TestNewBinaryTree_PostOrder(t *testing.T) {
	tree := NewBinaryTree()
	tree.AddNode(3)
	tree.AddNode(1)
	tree.AddNode(2)
	tree.AddNode(5)
	tree.AddNode(4)

	tree.PostOrder()
	//

}

func TestBinaryTree_PreOrder(t *testing.T) {
	tree := NewBinaryTree()
	tree.AddNode(3)
	tree.AddNode(1)
	tree.AddNode(2)
	tree.AddNode(5)
	tree.AddNode(4)
	//
	tree.PreOrder()
}

