package server

import (
	"log"
	"net/http"
)

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web-content/index.html")
}

// StartServer starts HTTP server that provides all static and dynamic data
func StartServer(port uint) {
	log.Printf("Starting server on port %d", port)
	/* http.Handle("/", http.FileServer(http.Dir("web-content/"))) */
	http.HandleFunc("/", indexPageHandler)
	http.ListenAndServe(":8080", nil)
}
