/*
 * routing.go
 * Copyright 2015 thewisenerd <thewisenerd@protonmail.com>
 *
 * Use of this source code is governed by a GNU GPL v2.0
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//TODO: write documentation.
	io.WriteString(w, http.StatusText(http.StatusOK))
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func muxInit() {
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello

	// trakker.go
	mux["/add"]  = addTracker
	mux["/list"] = listTrackers
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.Path]; ok {
		h(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, http.StatusText(http.StatusNotFound))
}
