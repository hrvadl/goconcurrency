package main

import (
	"net/http"
	"time"
)

const (
	timeout = time.Second * 2
)

func newServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: timeout,
		IdleTimeout:       timeout,
	}

	return srv
}
