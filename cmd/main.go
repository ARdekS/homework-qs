package main

import (
	"net/http"

	"github.com/ARdekS/homework-qs/internal/cache"
	"github.com/ARdekS/homework-qs/internal/handler"
	"github.com/ARdekS/homework-qs/internal/storage"
)

func main() {

	TreeStorage := storage.NewStorage()
	TreeCache := cache.NewCache()
	TreeStorage.InitTree()

	h := handler.NewHandler(TreeStorage, TreeCache)

	//handle requests
	http.HandleFunc("/", h.MainPage)
	http.HandleFunc("/addToCache", h.AddToCache)
	http.HandleFunc("/addToStorage", h.AddToStorage)
	http.HandleFunc("/renameItem", h.RenameItem)
	http.HandleFunc("/deleteItem", h.DeleteItem)
	http.HandleFunc("/newItem", h.NewItem)
	http.HandleFunc("/getCache", h.GetCache)
	http.HandleFunc("/getStorage", h.GetStorage)
	http.HandleFunc("/reset", h.Reset)
	http.ListenAndServe(":81", nil)
}
