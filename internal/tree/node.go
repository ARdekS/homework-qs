package tree

import "fmt"

// Глобальный счетчик для автоинкремента ID
var nodeID uint64

// Функция для генерации уникального ID
func generateID() uint64 {
	nodeID++
	return nodeID
}

type Node struct {
	Text     string `json:"text"`
	ParentID uint64 `json:"parent"`
	ID       uint64 `json:"id"`
	IsChild  bool   `json:"isChild,string"`
	Nodes    []*Node
}

func (n *Node) addChild(text *string) Node {
	id := generateID()
	var title string
	if text == nil {
		title = fmt.Sprintf("Task %d", id)
	} else {
		title = *text
	}

	c := Node{Text: title,
		ParentID: n.ID,
		ID:       id,
		Nodes:    make([]*Node, 0),
	}
	n.Nodes = append(n.Nodes, &c)
	return c
}

func (n *Node) Copy() *Node {
	return &Node{
		Text:     n.Text,
		ParentID: n.ParentID,
		ID:       n.ID,
		IsChild:  false,
		Nodes:    []*Node{},
	}
}

func (n *Node) Merge(m *Node) {
	// return &Node{
	// 	Text:   n.Text,
	// 	Parent: n.Parent,
	// 	ID:     n.ID,
	// 	IsChild:   false,
	// 	Nodes:  []*Node{},
	// }
	n.Text = m.Text
	n.IsChild = m.IsChild
}
