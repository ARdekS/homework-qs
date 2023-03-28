package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"

	"github.com/ARdekS/homework-qs/internal/cache"
	"github.com/ARdekS/homework-qs/internal/storage"
	"github.com/ARdekS/homework-qs/internal/tree"
)

type Handler struct {
	Storage *storage.Storage
	Cache   *cache.Cache
}

func NewHandler(s *storage.Storage, c *cache.Cache) *Handler {
	return &Handler{
		Storage: s,
		Cache:   c,
	}
}

func (h *Handler) getCacheTemplate() *template.Template {
	tmpl := template.New("cache").Funcs(
		template.FuncMap{
			"getNode": func(id int) *tree.Node {
				return h.Cache.GetNode(id)
			},
		},
	)
	return template.Must(tmpl.ParseFiles("./templates/cache.html", "./templates/item.html"))
}

func (h *Handler) getStorageTemplate() *template.Template {
	tmpl := template.New("storage").Funcs(
		template.FuncMap{
			"getNode": func(id int) *tree.Node {
				return h.Storage.GetNode(id)
			},
		},
	)
	return template.Must(tmpl.ParseFiles("./templates/storage.html", "./templates/item.html"))
}

func (h *Handler) RenameItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		var node tree.Node
		err = json.Unmarshal(body, &node)

		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		//Применение изменений в кэш
		h.Cache.EditItem(node)
		//формирование HTML и Ответа
		tmpl := h.getCacheTemplate()
		tmpl.ExecuteTemplate(w, "cache", h.Cache.ReturnItems())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		var node tree.Node
		err = json.Unmarshal(body, &node)

		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		//Применение изменений в кэш
		h.Cache.DeleteItem(node.ID)
		//формирование HTML и Ответа
		tmpl := h.getCacheTemplate()
		tmpl.ExecuteTemplate(w, "cache", h.Cache.ReturnItems())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
func (h *Handler) NewItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		var node tree.Node
		err = json.Unmarshal(body, &node)

		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		//Применение изменений в кэш
		h.Cache.NewItme(node)
		//формирование HTML и Ответа
		tmpl := h.getCacheTemplate()
		tmpl.ExecuteTemplate(w, "cache", h.Cache.ReturnItems())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) AddToCache(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		var node tree.Node
		err = json.Unmarshal(body, &node)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		//
		n := h.Storage.GetNode(node.ID)
		h.Cache.AddItem(n)
		//
		tmpl := h.getCacheTemplate()
		tmpl.ExecuteTemplate(w, "cache", h.Cache.ReturnItems())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
func (h *Handler) CacheToStorage(id int) {
	m := h.Cache.GetNode(id)
	h.Storage.AddToStorage(*m)
	for _, i := range m.Nodes {
		h.CacheToStorage(i)
	}
}

func (h *Handler) AddToStorage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		for _, n := range h.Cache.ReturnItems() {
			h.CacheToStorage(n)
		}

		tmpl := h.getStorageTemplate()
		tmpl.ExecuteTemplate(w, "storage", h.Storage.Head)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) GetCache(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := h.getCacheTemplate()
		tmpl.ExecuteTemplate(w, "cache", h.Cache.ReturnItems())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
func (h *Handler) Reset(w http.ResponseWriter, r *http.Request) {
	h.Cache = cache.NewCache()
	h.Storage = storage.NewStorage()
	h.Storage.InitTree()
	w.WriteHeader(200)
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/base.html"))
	tmpl.Execute(w, h)
}

func (h *Handler) GetStorage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := h.getStorageTemplate()
		tmpl.ExecuteTemplate(w, "storage", h.Storage.Head)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
