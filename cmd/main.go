package main

import (
	"encoding/json"
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
		var node tree.Node
		err = json.Unmarshal(body, &node)
		if err != nil {
			// Обработка ошибки распаковки JSON
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		treeMainPage.TreeCache.EditItem(&node)
		tmpl := template.Must(template.ParseFiles("./templates/cache.html", "./templates/item.html"))
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
		treeMainPage.TreeCache.AddItem(&node)
		tmpl := template.Must(template.ParseFiles("./templates/cache.html", "./templates/item.html"))
		tmpl.Execute(w, treeMainPage.TreeCache.ReturnItems())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

}
func addToStorage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		for _, n := range treeMainPage.TreeCache.Cache {
			treeMainPage.TreeStorage.AddToStorage(n)
		}
		tmpl := template.Must(template.ParseFiles("./templates/storage.html", "./templates/item.html"))
		tmpl.Execute(w, treeMainPage.TreeStorage.Head)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/base.html", "./templates/item.html"))
	tmpl.Execute(w, treeMainPage)
}

func main() {
	treeMainPage.TreeCache = cache.NewCache()
	treeMainPage.TreeStorage = tree.InitTree()

	// Обходим дерево
	// tree.Traverse(treeStorage.Head)
	//handle requests
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/addToCache", addToCache)
	http.HandleFunc("/addToStorage", addToStorage)
	http.HandleFunc("/renameItem", renameItem)
	http.ListenAndServe(":81", nil)
}
