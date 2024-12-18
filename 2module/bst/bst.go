package main

import "fmt"

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

type Queue interface {
	Size() int
	Enqueue(e int)
	Dequeue() (int, error)
	Front() (int, error)
	// IsEmpty() bool
}

type Node struct {
	val  *BSTNode
	next *Node
}

type LinkedListQueue struct {
	front    *Node
	rear     *Node
	inserted int
}

func (queue *LinkedListQueue) Front() (int, error) {
	if queue.inserted == 0 {
		return -9999, nil
	}
	return queue.front.val.val, nil
}

func (queue *LinkedListQueue) Size() int {
	return queue.inserted
}

func (queue *LinkedListQueue) Enqueue(node *BSTNode) {
	newNode := &Node{val: node}
	if queue.inserted == 0 {
		queue.front = newNode
		queue.rear = newNode
	} else {
		queue.rear.next = newNode
		queue.rear = newNode
	}
	queue.inserted++
}

func (queue *LinkedListQueue) Dequeue() (int, error) {
	if queue.inserted == 0 {
		return -1, nil
	}
	if queue.inserted == 1 {
		val := queue.front.val.val
		queue.front = nil
		queue.rear = nil
		queue.inserted--
		return val, nil
	}

	val := queue.front.val.val
	queue.front = queue.front.next

	queue.inserted--
	return val, nil
}

func createNode(val int) *BSTNode {
	return &BSTNode{val: val}
}

func (node *BSTNode) Add(val int) {
	if node.val > val {
		if node.left != nil {
			node.left.Add(val)
		} else {
			node.left = createNode(val)
		}
	} else {
		if node.right != nil {
			node.right.Add(val)

		} else {
			node.right = createNode(val)
		}
	}
}

func (node *BSTNode) Search(val int) bool {
	if val == node.val {
		return true
	}
	if val < node.val {
		if node.left == nil {
			return false
		} else {
			return node.left.Search(val)
		}
	} else {
		if node.right == nil {
			return false
		} else {
			return node.right.Search(val)
		}
	}
}

func (node *BSTNode) Height() int {
	if node == nil {
		return -1
	} else {
		return 1 + max(node.left.Height(), node.right.Height())
	}
}

func (node *BSTNode) Min() int {
	if node.left == nil {
		return node.val
	} else {
		return node.left.Min()
	}
}

func (node *BSTNode) Max() int {
	if node.right == nil {
		return node.val
	} else {
		return node.right.Max()
	}
}

func (node *BSTNode) Imprimir() {
	if node.left != nil {
		node.left.Imprimir()
	}

	fmt.Print(node.val)
	fmt.Print(", ")

	if node.right != nil {
		node.right.Imprimir()
	}

}

func (node *BSTNode) navNivel() {
	queue := &LinkedListQueue{front: nil, rear: nil, inserted: 0}

	queue.Enqueue(node)

	for queue.inserted > 0 {
		elem := queue.front.val
		fmt.Print(elem.val)
		fmt.Print(", ")

		if elem.left != nil {
			queue.Enqueue(elem.left)

		}
		if elem.right != nil {
			queue.Enqueue(elem.right)
		}

		queue.Dequeue()
	}
}

func (node *BSTNode) navPreOrdem() {
	if node != nil {
		fmt.Print(node.val)
		fmt.Print(", ")

		node.left.navPreOrdem()

		node.right.navPreOrdem()
	}
}

func (node *BSTNode) navEmOrdem() {
	if node != nil {
		node.left.navEmOrdem()

		fmt.Print(node.val)
		fmt.Print(", ")

		node.right.navEmOrdem()
	}
}

func (node *BSTNode) navPosOrdem() {
	if node != nil {
		node.left.navPosOrdem()
		node.right.navPosOrdem()

		fmt.Print(node.val)
		fmt.Print(", ")
	}
}

func (node *BSTNode) remove(val int) *BSTNode {
	if node != nil {
		if node.val > val {
			node.left = node.left.remove(val)
		} else if node.val < val {
			node.right = node.right.remove(val)
		} else { //found
			if node.right == nil && node.left == nil {
				return nil
			} else if node.left == nil {
				return node.right
			} else if node.right == nil {
				return node.left
			} else {
				max_esq := node.left.Max()
				node.val = max_esq
				node.left = node.left.remove(max_esq)
			}
		}
	}
	return node
}

func main() {
	root := createNode(12)

	root.Add(8)
	root.Add(14)
	root.Add(4)
	root.Add(10)
	root.Add(2)
	root.Add(6)
	root.Add(18)
	root.Add(16)
	root.Add(20)
	// root.Add(58)
	// root.Add(59)

	// root.Imprimir()

	// fmt.Println(root.Search(3))
	// fmt.Println(root.Search(22))

	// fmt.Println(root.Max())
	// fmt.Println(root.Min())

	// fmt.Println(root.Height())

	// root.navNivel()

	// root.navPreOrdem()
	// fmt.Println()
	// root.navEmOrdem()
	// fmt.Println()
	// root.navPosOrdem()
	// fmt.Println()

	root.Imprimir()
	fmt.Println()

	root.remove(6)
	root.Imprimir()
	fmt.Println()

	root.remove(4)
	root.Imprimir()
	fmt.Println()

	root.remove(12)
	root.Imprimir()
	fmt.Println()

	root.remove(18)
	root.Imprimir()
	fmt.Println()
}
