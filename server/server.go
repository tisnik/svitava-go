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

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web-content/index.html")
}

func newFractalPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web-content/new_fractal.html")
}

func galleryPageHandler(w http.ResponseWriter, r *http.Request) {
}

func settingsPageHandler(w http.ResponseWriter, r *http.Request) {
}

func staticImageHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.String()
	fileName := strings.TrimPrefix(imageName, "/image/")
	http.ServeFile(w, r, "web-content/images/"+fileName)
}

// StartServer starts HTTP server that provides all static and dynamic data
func StartServer(port uint) {
	log.Printf("Starting server on port %d", port)
	/* http.Handle("/", http.FileServer(http.Dir("web-content/"))) */

	http.HandleFunc("/", indexPageHandler)
	http.HandleFunc("/new-fractal", newFractalPageHandler)
	http.HandleFunc("/gallery", galleryPageHandler)
	http.HandleFunc("/settings", settingsPageHandler)
	//http.HandleFunc("/image/", staticImageHandler)
	imageServer := http.FileServer(http.Dir("web-content/images/"))
	http.Handle("/image/", http.StripPrefix("/image", imageServer))

	http.ListenAndServe(":8080", nil)
}
