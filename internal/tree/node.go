package tree

import "fmt"

// // Глобальный счетчик для автоинкремента ID

func (n *Node) addChild() Node {
	id := generateID()
	title := fmt.Sprintf("Task %d", id)

	c := Node{
		Text:     title,
		ParentID: n.ID,
		ID:       id,
		Nodes:    make([]int, 0),
	}
	n.Nodes = append(n.Nodes, c.ID)
	return c
}

func (n *Node) Copy() *Node {
	return &Node{
		Text:     n.Text,
		ParentID: n.ParentID,
		ID:       n.ID,
		Nodes:    []int{},
	}
}
