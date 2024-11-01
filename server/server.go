//
//  (C) Copyright 2019, 2020, 2021, 2022, 2023, 2024  Pavel Tisnovsky
//
//  All rights reserved. This program and the accompanying materials
//  are made available under the terms of the Eclipse Public License v1.0
//  which accompanies this distribution, and is available at
//  http://www.eclipse.org/legal/epl-v10.html
//
//  Contributors:
//      Pavel Tisnovsky
//

package server

import (
	"log"
	"net/http"
	"strings"
)

type Server interface {
	Serve()
}

type HTTPServer struct {
	port uint
}

func NewHTTPServer(port uint) Server {
	return HTTPServer{
		port: port,
	}
}

func (s HTTPServer) indexPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web-content/index.html")
}

func (s HTTPServer) newFractalPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web-content/new_fractal.html")
}

func (s HTTPServer) galleryPageHandler(w http.ResponseWriter, r *http.Request) {
}

func (s HTTPServer) settingsPageHandler(w http.ResponseWriter, r *http.Request) {
}

func (s HTTPServer) staticImageHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.String()
	fileName := strings.TrimPrefix(imageName, "/image/")
	http.ServeFile(w, r, "web-content/images/"+fileName)
}

// Serve starts HTTP server that provides all static and dynamic data
func (s HTTPServer) Serve() {
	log.Printf("Starting server on port %d", s.port)
	/* http.Handle("/", http.FileServer(http.Dir("web-content/"))) */

	http.HandleFunc("/", s.indexPageHandler)
	http.HandleFunc("/new-fractal", s.newFractalPageHandler)
	http.HandleFunc("/gallery", s.galleryPageHandler)
	http.HandleFunc("/settings", s.settingsPageHandler)
	//http.HandleFunc("/image/", staticImageHandler)
	imageServer := http.FileServer(http.Dir("web-content/images/"))
	http.Handle("/image/", http.StripPrefix("/image", imageServer))

	http.ListenAndServe(":8080", nil)
}
