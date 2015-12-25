/*
 * list.go
 * Copyright 2015 thewisenerd <thewisenerd@protonmail.com>
 *
 * Use of this source code is governed by a GNU GPL v2.0
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	// _ "fmt"
	"io"
	"net/http"
)

var trakkers []Tracker

func addTrackertoList(t Tracker) (e bool) {

	if trakkers == nil {
		trakkers = make([]Tracker, 1)
		trakkers[0] = t
		return true
	}

	for _, v := range trakkers {
		if t.name == v.name && t.port == t.port {
			return false
		}
	}

	trakkers = append(trakkers, t)

	return true
}

func ListTrackers(w http.ResponseWriter, r *http.Request) {
	for _, v := range trakkers {
		io.WriteString(w, v.fullname);
		io.WriteString(w, "\n\n");
	}
}