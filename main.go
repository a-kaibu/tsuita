package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed tsuita.png index.html
var content embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := content.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	})

	port := ":8080"
	fmt.Printf("🚀 Server running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
