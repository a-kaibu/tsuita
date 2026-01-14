package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed tsuita.png index.html
var content embed.FS

func main() {
	fileServer := http.FileServer(http.FS(content))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			data, err := fs.ReadFile(content, "index.html")
			if err != nil {
				http.Error(w, "Page not found", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(data)
			return
		}
		fileServer.ServeHTTP(w, r)
	})

	port := ":8080"
	fmt.Printf("🚀 Server running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
