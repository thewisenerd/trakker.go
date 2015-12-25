/*
 * trakker.go
 * Copyright 2015 thewisenerd <thewisenerd@protonmail.com>
 *
 * Use of this source code is governed by a GNU GPL v2.0
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	//_"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var trackerRegex *regexp.Regexp

type Tracker struct {
	protocol string
	name     string
	port     int64
	fullname string
}

func trakkerAddURL(s string) (e bool) {
	res := trackerRegex.FindAllStringSubmatch(s, -1)
	m   := res[0]

	var trakker Tracker

	trakker.protocol = m[2]

	if (m[3] == m[4]) { // url:port/abc
		// set port
		trakker.port, _ = strconv.ParseInt(m[5], 0, 64)

		trakker.name = m[4]
		trakker.fullname = m[0]

		return addTrackertoList(trakker)
	} else if (m[3] == m[8]) { // url
		trakker.port = 80 // TODO: verify this for UDP?

		// TODO: there should be a better way to do this?
		trakker.name = m[8][0:len(m[8])-len(m[9] + m[10])]
		trakker.fullname = m[0]

		return addTrackertoList(trakker)
	}

	return false
}

func listTrackers(w http.ResponseWriter, r *http.Request) {
	ListTrackers(w, r)
}

func addTracker(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	// url param not set
	if len(q) == 0 {
		// malformed request
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "missing url param")
		return
	}

	url := r.URL.Query().Get("url")

	// url param empty
	if url == "" {
		// malformed request
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "missing url param")
		return
	}

	log.Println("url: " + url)

	match := trackerRegex.MatchString(url)

	if match == true {
		ret := trakkerAddURL(url)
		if ret == true {
			io.WriteString(w, "OK")
		} else {
			io.WriteString(w, "fail")
		}
	} else {
		// malformed url
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad url param")
	}
}

func trakkerInit() {

/*
 * https://regex101.com/r/oF6eO5/1
 *
 * start:
 * ((udp|http):\/{2})
 *
 * a = url:port or url:port/ or url:port/abc
 * ([^:/]+:([0-9]+)(\/)?(.+)?)
 *
 * b = url or url/ or url/abc
 * ([^:/]+(\/)?(a-zA-Z0-9+)?)
 *
 * regex = start(a|b)
 */
	trackerRegex = regexp.MustCompile("((udp|http):/{2})(([^:/]+:([0-9]+)(/)?(.+)?)|([^:/]+(/)?([a-zA-Z0-9]+)?))")
}