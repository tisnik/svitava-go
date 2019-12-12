package server

import (
	"log"
	"net/http"
)

// StartServer starts HTTP server that provides all static and dynamic data
func StartServer(port uint) {
	log.Printf("Starting server on port %d", port)
	http.Handle("/", http.FileServer(http.Dir("web-content/")))
	/* http.HandleFunc("/", someHandler) */
	http.ListenAndServe(":8080", nil)
}
