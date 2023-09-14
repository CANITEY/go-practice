package main

import (
	"net/http"
	"github.com/canitey/simpleTaskManager/internal/server"
)

func main() {
	http.HandleFunc("/", server.Index)
	http.HandleFunc("/toggle/", server.Toggle)
	http.HandleFunc("/add/", server.Add)
	http.HandleFunc("/delete/", server.Delete)
	http.ListenAndServe(":8888", nil)
}
