package main

import (
	"log/slog"
	"net"
)

func main() {
	addr := net.JoinHostPort("localhost", "3000")
	srv := newServer(addr)
	slog.Info("Starting the server", slog.String("addr", addr))
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Failed to listen", slog.Any("err", err))
	}
}
