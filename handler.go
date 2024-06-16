package main

import (
	"net/http"

	"gitlab.com/golang-commonmark/mdurl"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{Query: mdurl.Decode(r.URL.RawQuery), Body: r.FormValue("body")}
	index.Execute(w, data)
}

func addHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Powered-By", telerWAF)
		w.Header().Set("Server", server)

		next.ServeHTTP(w, r)
	})
}
