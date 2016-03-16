package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	addr = "127.0.0.1:8090"
)

var (
	indexTemplate = template.Must(template.ParseFiles("index.html"))
	tasks         = []Task{
		Task{1, "Learn golang", true},
		Task{2, "Write webapp", true},
		Task{3, "???", false},
		Task{4, "Profit! $$$", false},
	}
)

type Index struct {
	Tasks []Task
	Addr  string
}

type Task struct {
	Id   uint64
	Name string
	Done bool
}

func index(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, Index{tasks, addr})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/ws", serveWs)
	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(static)

	go h.run()
	logger := handlers.LoggingHandler(os.Stdout, r)
	log.Println("Starting server on", addr)
	log.Fatal(http.ListenAndServe(addr, logger))
}
