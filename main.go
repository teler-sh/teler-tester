package main

import (
	"net"
	"os"

	"net/http"

	"github.com/teler-sh/teler-waf"
)

func init() {
	portEnv := os.Getenv("PORT")
	if portEnv != "" {
		port = portEnv
	}

	server = telerWAF
	if telerWAFVersion != "" {
		server += "/" + telerWAFVersion
	}
}

func main() {
	waf := teler.New()
	app := waf.Handler(http.HandlerFunc(myHandler))

	http.Handle("/static/", addHeaders(http.StripPrefix("/static/", static)))
	http.Handle("/", addHeaders(app))
	http.Handle("/healthz", addHeaders(http.HandlerFunc(healthz)))

	go func() {
		println("Listening on " + port)
	}()

	http.ListenAndServe(net.JoinHostPort("0.0.0.0", port), nil)
}
