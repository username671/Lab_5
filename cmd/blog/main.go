package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = "3000"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", GetHome)
	mux.HandleFunc("/post", GetPost)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Listen on port", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
