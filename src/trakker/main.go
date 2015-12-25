/*
 * main.go
 * Copyright 2015 thewisenerd <thewisenerd@protonmail.com>
 *
 * Use of this source code is governed by a GNU GPL v2.0
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("initializng trakker\n")

	// initialize mux
	muxInit()

	// initialize trakker
	trakkerInit()

	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	log.Fatal(server.ListenAndServe())
}
