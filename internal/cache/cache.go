package cache

import (
	"sort"

	"github.com/ARdekS/homework-qs/internal/tree"
)

// type ID struct {
// 	int `json:"id"`
// }

type Cache struct {
	Cache map[int]*Node
}
type Node struct {
	tree.Node
	IsLinked bool `json:"isLinked,string"`
}

func TreeToCache(n *tree.Node) *Node {
	return &Node{
		Node: tree.Node{
			Text:     n.Text,
			ParentID: n.ParentID,
			ID:       n.ID,
			Nodes:    []int{},
		},
		IsLinked: false,
	}
}
func NewCache() *Cache {
	return &Cache{
		Cache: make(map[int]*Node, 0),
	}
}
func (c *Cache) GetNode(id int) *Node {
	return c.Cache[id]
}

func (c *Cache) AddItem(n *tree.Node) {
	if _, ok := c.Cache[n.ID]; !ok {
		m := TreeToCache(n)
		c.Cache[m.ID] = m
		//если n это чей либо родитель то мы должны добавить детей и скрыть их
		for _, i := range n.Nodes {
			if child, ok := c.Cache[i]; ok {
				m.Nodes = append(m.Nodes, child.ID)
				child.IsLinked = true
			}
		}
		// for _, i := range c.Cache {
		// 	if i.ParentID == n.ID {
		// 		n.Nodes = append(n.Nodes, i.ID)
		// 		i.IsLinked = true
		// 	}
		// }
		//если родитель для n уже есть то мы должны добавить ему ребенка n
		if p, ok := c.Cache[m.ParentID]; ok {
			m.IsLinked = true
			p.Nodes = append(p.Nodes, m.ID)
		}
	}
}
func (c *Cache) EditItem(m Node) {
	if n, ok := c.Cache[m.ID]; ok {
		n.Text = m.Text
	}
}

func (c *Cache) ReturnItems() []int {
	var result []int = make([]int, 0, 1)
	for k, v := range c.Cache {
		if !v.IsLinked {
			result = append(result, k)
		}
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
