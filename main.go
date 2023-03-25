package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// type Todo struct {
// 	Title string
// 	Done  bool
// }

type Node struct {
	Text   string `json:"text"`
	Parent int    `json:"parent"`
	ID     int    `json:"id"`
	Index  int
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
	return &data
}

type Cache struct {
	Storage map[int]*Node
}

type ID struct {
	ID int `json:"id"`
}

func NewCache() *Cache {
	return &Cache{Storage: make(map[int]*Node, 0)}
}
func (c *Cache) addItem(n *Node) {
	c.Storage[n.ID] = n
	if p, ok := c.Storage[n.Parent]; ok {
		p.Nodes = append(p.Nodes, n)
	}
}

var cache = NewCache()

func addRequest(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method == http.MethodPost {
		// Читаем тело (body) запроса
		body, err := ioutil.ReadAll(r.Body)
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
		fmt.Println(cache)
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
	tmpl := template.Must(template.ParseFiles("cache.html"))
	tmpl.Execute(w, cache)
}

var data *Node

type foo struct {
	Storage *Node
	Cache   *Cache
}

func main() {
	data = initTree()

	// Обходим дерево
	traverse(data)
	tmpl := template.Must(template.ParseFiles("base.html", "item.html", "cache.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	http.HandleFunc("/method", addRequest)

	http.ListenAndServe(":88", nil)
}
