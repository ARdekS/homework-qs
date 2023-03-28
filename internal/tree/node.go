package tree

import "fmt"

// // Глобальный счетчик для автоинкремента ID
var nodeID int

func ResetID() {
	nodeID = 0
}

// Функция для генерации уникального ID
func GenerateID() int {
	nodeID++
	return nodeID
}

type Node struct {
	Text      string `json:"text"`
	ParentID  int    `json:"parent"`
	ID        int    `json:"id"`
	IsDeleted bool   `json:"isDeleted,string"`
	Nodes     []int
}

func (n *Node) AddChild() Node {
	id := GenerateID()
	title := fmt.Sprintf("Task %d", id)

	c := Node{
		Text:     title,
		ParentID: n.ID,
		ID:       id,
		Nodes:    make([]int, 0),
	}
	return c
}
func (n *Node) Edit(m Node) {
	n.Text = m.Text
}

func (n *Node) Copy() Node {
	return Node{
		Text:      n.Text,
		ParentID:  n.ParentID,
		ID:        n.ID,
		IsDeleted: n.IsDeleted,
		Nodes:     make([]int, 0),
	}
}
