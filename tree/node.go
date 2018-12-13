package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, "  ")
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node *Node) SetValue(value int){
	if node == nil {
		fmt.Println("Setting Value to nil" +
			"node, Ignored")
		return
	}
	node.Value = value
}


