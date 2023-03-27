package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/ARdekS/homework-qs/internal/cache"
	"github.com/ARdekS/homework-qs/internal/tree"
)

type MainPage struct {
	TreeCache   *cache.Cache
	TreeStorage *tree.Storage
}

var treeMainPage MainPage

func renameItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		// Распарсим JSON в структуру MyStruct
		var node cache.Node
		err = json.Unmarshal(body, &node)
		if err != nil {
			// Обработка ошибки распаковки JSON
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		treeMainPage.TreeCache.EditItem(node)
		tmpl := template.New("cache").Funcs(
			template.FuncMap{
				"getNode": func(id int) *cache.Node {
					return treeMainPage.TreeCache.GetNode(id)
				},
			},
		)
		tmpl = template.Must(tmpl.ParseFiles("./templates/cache.html", "./templates/item.html"))
		tmpl.Execute(w, treeMainPage.TreeCache.ReturnItems())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
func addToCache(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		// Распарсим JSON в структуру MyStruct
		var node tree.Node
		err = json.Unmarshal(body, &node)
		if err != nil {
			// Обработка ошибки распаковки JSON
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		n := treeMainPage.TreeStorage.GetNode(node.ID)
		treeMainPage.TreeCache.AddItem(n)
		tmpl := template.New("cache").Funcs(
			template.FuncMap{
				"getNode": func(id int) *cache.Node {
					return treeMainPage.TreeCache.GetNode(id)
				},
			},
		)
		tmpl = template.Must(tmpl.ParseFiles("./templates/cache.html", "./templates/item.html"))
		tmpl.Execute(w, treeMainPage.TreeCache.ReturnItems())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

}
func addToStorage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		for _, n := range treeMainPage.TreeCache.Cache {
			treeMainPage.TreeStorage.AddToStorage(n.Node)
		}
		tmpl := template.New("storage").Funcs(
			template.FuncMap{
				"getNode": func(id int) *tree.Node {
					return treeMainPage.TreeStorage.GetNode(id)
				},
			},
		)
		tmpl = template.Must(tmpl.ParseFiles("./templates/storage.html", "./templates/item.html"))
		tmpl.Execute(w, treeMainPage.TreeStorage.Head)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func getStorage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tmpl := template.New("storage").Funcs(
			template.FuncMap{
				"getNode": func(id int) *tree.Node {
					return treeMainPage.TreeStorage.GetNode(id)
				},
			},
		)
		tmpl = template.Must(tmpl.ParseFiles("./templates/storage.html", "./templates/item.html"))
		tmpl.Execute(w, treeMainPage.TreeStorage.Head)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func getCache(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tmpl := template.New("cache").Funcs(
			template.FuncMap{
				"getNode": func(id int) *cache.Node {
					return treeMainPage.TreeCache.GetNode(id)
				},
			},
		)
		tmpl = template.Must(tmpl.ParseFiles("./templates/cache.html", "./templates/item.html"))
		tmpl.Execute(w, treeMainPage.TreeCache.ReturnItems())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.New("main").ParseFiles("./templates/main.html", "./templates/node.html")
	t := template.New("base").Funcs(
		template.FuncMap{
			"getNode": func(id int) *tree.Node {
				return treeMainPage.TreeStorage.GetNode(id)
			},
		},
	)
	t, _ = t.ParseFiles("./templates/base.html", "./templates/item.html")

	// t, _ := template.New("main").ParseFiles("./templates/main.html")
	// if err != nil {
	// 	// обработка ошибок
	// }
	// t2 := template.New("node")
	// t = template.Must(t.ParseFiles())

	// data := struct {
	// 	Storage *tree.Storage
	// 	Cache   *cache.Cache
	// }{
	// 	treeMainPage.TreeStorage,
	// 	treeMainPage.TreeCache,
	// }
	t.Execute(w, treeMainPage)
	// tmpl := template.Must(template.New("node").ParseFiles("./templates/node.html"))
	// tmpl := template.Must(template.New("base").Funcs(template.FuncMap{
	// 	"getNode": func(id int) *tree.Node {
	// 		return treeMainPage.TreeStorage.GetNode(id)
	// 	},
	// }).ParseFiles("./templates/base.html", "./templates/item.html"))
	// tmpl = template.Must(tmpl.ParseFiles("./templates/base.html"))
	// tmpl.ExecuteTemplate(w, "nodex", treeMainPage)
	// tmpl.Execute(w, treeMainPage)
}

func main() {
	treeMainPage.TreeCache = cache.NewCache()
	treeMainPage.TreeStorage = tree.InitTree()
	fmt.Println(treeMainPage.TreeStorage.Head)
	// Обходим дерево
	treeMainPage.TreeStorage.Traverse(treeMainPage.TreeStorage.Head)
	//handle requests
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/addToCache", addToCache)
	http.HandleFunc("/addToStorage", addToStorage)
	http.HandleFunc("/renameItem", renameItem)
	http.HandleFunc("/getCache", getCache)
	http.HandleFunc("/getStorage", getStorage)
	http.ListenAndServe(":81", nil)
}
