package main

import (
	"bytes"
	"log/slog"
	"net/http"
	"sync"
)

var pool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func handler(w http.ResponseWriter, _ *http.Request) {
	buf, ok := pool.Get().(*bytes.Buffer)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buf.Reset()
	defer pool.Put(buf)

	if _, err := buf.WriteString("hello"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := buf.WriteString(" world!"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		slog.Error("Failed to write response", slog.Any("err", err))
	}
}
