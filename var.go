package main

import "net/http"

var (
	static = http.FileServer(http.Dir("static"))
	port   = "3000"

	telerWAF = "teler-waf"
	server   = telerWAF

	telerWAFVersion string
)
