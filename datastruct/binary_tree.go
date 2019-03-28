package datastruct

import "fmt"

/*二叉树*/
type BinaryTree struct {
	root *TreeNode
}

type TreeNode struct {
	left *TreeNode
	right *TreeNode
	data int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) AddNode(data int)  {
	node := &TreeNode{
		data:data,
	}
	if b.root == nil {
		b.root = node
		return
	}

	if b.root.data > data {
		if b.root.left == nil {
			b.root.left = node
			return
		}
		b.root.left.AddNode(node)
	} else {
		if b.root.right == nil {
			b.root.right = node
			return
		}
		b.root.right.AddNode(node)
	}

}

func (t *TreeNode) AddNode(node* TreeNode)  {
	if t.data > node.data {
		if t.left == nil {
			t.left = node
			return
		}
		t.left.AddNode(node)
	} else {
		if t.right == nil {
			t.right = node
			return
		}
		t.right.AddNode(node)
	}
}

/*前序遍历*/
func (b *BinaryTree) PreOrder()  {
	if b.root == nil {
		return
	}
	fmt.Println(b.root.data)
	b.root.left.PreOrder()
	b.root.right.PreOrder()
}

func (t *TreeNode) PreOrder()  {
	if t == nil {
		return
	}
	fmt.Println(t.data)
	if t.left != nil {
		t.left.PreOrder()
	}
	if t.right != nil {
		t.right.PreOrder()
	}
}

/*后续遍历*/
func (b *BinaryTree) PostOrder()  {
	if b.root == nil {
		return
	}
	b.root.left.PostOrder()
	b.root.right.PostOrder()
	fmt.Println(b.root.data)
}

func (t *TreeNode) PostOrder()  {
	if t == nil {
		return
	}

	if t.left != nil {
		t.left.PostOrder()
	}
	if t.right != nil {
		t.right.PostOrder()
	}
	fmt.Println(t.data)
}

/*中序遍历*/
func (b *BinaryTree) MidOrder()  {
	if b.root == nil {
		return
	}
	b.root.left.MidOrder()
	fmt.Println(b.root.data)
	b.root.right.MidOrder()

}

func (t *TreeNode) MidOrder()  {
	if t == nil {
		return
	}

	if t.left != nil {
		t.left.MidOrder()
	}
	if t.right != nil {
		t.right.MidOrder()
	}
	fmt.Println(t.data)
}

