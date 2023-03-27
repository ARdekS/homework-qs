package tree

import (
	"fmt"
)

var nodeID int

// Функция для генерации уникального ID
func generateID() int {
	nodeID++
	return nodeID
}

type Node struct {
	Text     string `json:"text"`
	ParentID int    `json:"parent"`
	ID       int    `json:"id"`
	Nodes    []int
}

type Storage struct {
	Storage map[int]*Node
	Head    int
}

func NewStorage() *Storage {
	return &Storage{Storage: make(map[int]*Node, 0), Head: 0}
}

func (s *Storage) AddToStorage(n Node) {
	//add head to storage
	if n.ParentID == 0 {
		s.Storage[n.ID] = &n
		s.Head = n.ID
		return
	}
	//Node already in Storage
	if i, ok := s.Storage[n.ID]; ok {
		i.Text = n.Text
		for _, childID := range n.Nodes {
			if _, ok = s.Storage[childID]; !ok {
				i.Nodes = append(i.Nodes, childID)
			}
		}
	} else {
		s.Storage[n.ID] = &n
		// if p, ok := s.Storage[n.ParentID]; ok {
		// 	p.Nodes = append(p.Nodes, m.ID)
		// }
	}

	// for _, child := range n.Nodes {
	// 	s.AddToStorage(child)
	// }

	// s.Storage[n.ID] = n

}

func (s *Storage) Traverse(i int) {
	m := s.Storage[i]
	fmt.Println(m.Text)
	for _, child := range m.Nodes {
		s.Traverse(child)
	}
}

func InitTree() *Storage {
	storage := NewStorage()

	data := Node{
		Text: "My TODO list", ID: generateID(), ParentID: 0, Nodes: make([]int, 0)}
	storage.AddToStorage(data)

	for i := 0; i < 3; i++ {
		storage.AddToStorage(storage.Storage[1].addChild())
	}
	for i := 0; i < 3; i++ {
		// b := data.Nodes[1].addChild()
		storage.AddToStorage(storage.Storage[2].addChild())
		// c := data.Nodes[2].addChild()
		storage.AddToStorage(storage.Storage[3].addChild())
	}
	storage.AddToStorage(storage.Storage[6].addChild())
	storage.AddToStorage(storage.Storage[7].addChild())
	storage.AddToStorage(storage.Storage[8].addChild())

	return storage
}

// func (s *Storage) ReturnItems(int) []*Node {
// 	var result []*Node = make([]*Node, 0, 1)
// 	for _, v := range s.Storage {
// 		result = append(result, v)

// 	}
// 	sort.SliceStable(result, func(i, j int) bool {
// 		return result[i].ID < result[j].ID
// 	})
// 	return result

// 	m := s.Storage[i]
// 	fmt.Println(m.Text)
// 	for _, child := range m.Nodes {
// 		s.Traverse(child)
// 	}
// }

func (s *Storage) GetNode(id int) *Node {
	return s.Storage[id]
}
