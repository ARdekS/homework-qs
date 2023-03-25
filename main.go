package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sort"
)

// type Todo struct {
// 	Title string
// 	Done  bool
// }

type Node struct {
	Text   string `json:"text"`
	Parent int    `json:"parent"`
	ID     int    `json:"id"`
	Hide   bool
	Nodes  []*Node
}

// Глобальный счетчик для автоинкремента ID
var nodeID int

// Функция для генерации уникального ID
func generateID() int {
	nodeID++
	return nodeID
}

func (t *Node) addNew() {
	id := generateID()
	n := &Node{Text: fmt.Sprintf("Task %d", id), Parent: t.ID, ID: id, Nodes: make([]*Node, 0)}
	t.Nodes = append(t.Nodes, n)
}

func traverse(node *Node) {
	fmt.Println(node.Text)
	for _, child := range node.Nodes {
		traverse(child)
	}
}
func initTree() *Node {
	data := Node{
		Text: "My TODO list", ID: 0}
	for i := 0; i < 3; i++ {
		data.addNew()
	}
	for i := 0; i < 3; i++ {
		data.Nodes[1].addNew()
		data.Nodes[2].addNew()
	}
	data.Nodes[2].addNew()
	data.Nodes[1].Nodes[2].addNew()
	data.Nodes[1].Nodes[1].addNew()
	data.Nodes[1].Nodes[1].Nodes[0].addNew()
	return &data
}

type Cache struct {
	Cache map[int]*Node
}

type ID struct {
	ID int `json:"id"`
}

func NewCache() *Cache {
	return &Cache{Cache: make(map[int]*Node, 0)}
}
func (c *Cache) addItem(n *Node) {
	if _, ok := c.Cache[n.ID]; !ok {
		//если n это чей либо родитель то мы должны добавить детей и скрыть их
		for _, i := range c.Cache {
			if i.Parent == n.ID {
				n.Nodes = append(n.Nodes, i)
				i.Hide = true
			}
		}
		//если родитель для n уже есть то мы должны добавить ему ребенка n
		if p, ok := c.Cache[n.Parent]; ok {
			n.Hide = true
			p.Nodes = append(p.Nodes, n)
		}

		c.Cache[n.ID] = n

	}
}
func (c *Cache) returnItems() []*Node {

	var result []*Node = make([]*Node, 0, 1)
	for _, v := range c.Cache {
		if !v.Hide {
			result = append(result, v)
		}
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

var cache = NewCache()

func addRequest(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method == http.MethodPost {
		// Читаем тело (body) запроса
		body, err := io.ReadAll(r.Body)
		if err != nil {

			// Обработка ошибки чтения тела (body) запроса
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		// Распарсим JSON в структуру MyStruct
		var node Node
		err = json.Unmarshal(body, &node)
		// fmt.Println(node)
		cache.addItem(&node)
		fmt.Println(cache.returnItems())
		if err != nil {
			// Обработка ошибки распаковки JSON
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
	} else {
		// Обработка других методов запроса
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	tmpl := template.Must(template.ParseFiles("cache.html", "item.html"))

	tmpl.Execute(w, cache.returnItems())
}

var data *Node

// type foo struct {
// 	Storage *Node
// 	Cache   *Cache
// }

func main() {
	data = initTree()

	// Обходим дерево
	traverse(data)
	tmpl := template.Must(template.ParseFiles("base.html", "item.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	http.HandleFunc("/method", addRequest)

	http.ListenAndServe(":88", nil)
}
