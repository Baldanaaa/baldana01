package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет из Snippetbox"))
}
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("отображение заметки..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("форма для создания новый заметки..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
